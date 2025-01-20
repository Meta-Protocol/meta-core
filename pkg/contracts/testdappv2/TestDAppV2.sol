// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

struct RevertOptions {
    address revertAddress;
    bool callOnRevert;
    address abortAddress;
    bytes revertMessage;
    uint256 onRevertGasLimit;
}

interface IGatewayZEVM {
    function withdraw(
        bytes memory receiver,
        uint256 amount,
        address zrc20,
        RevertOptions calldata revertOptions
    ) external;

    function call(
        bytes memory receiver,
        address zrc20,
        bytes calldata message,
        uint256 gasLimit,
        RevertOptions calldata revertOptions
    ) external;
}

interface IGatewayEVM {
    function deposit(address receiver, RevertOptions calldata revertOptions) external payable;
    function depositAndCall(
        address receiver,
        bytes calldata payload,
        RevertOptions calldata revertOptions
    )
    external
    payable;
    function call(address receiver, bytes calldata payload, RevertOptions calldata revertOptions) external;
}

interface IZRC20 {
    function approve(address spender, uint256 amount) external returns (bool);
    function withdrawGasFee() external view returns (address, uint256);
}

interface IERC20 {
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
}

contract TestDAppV2 {
    // used to simulate gas consumption
    uint256[] private storageArray;

    string public constant NO_MESSAGE_CALL = "called with no message";

    // define if the chain is ZetaChain
    bool immutable public isZetaChain;

    // address of the gateway
    address immutable public gateway;

    struct zContext {
        bytes origin;
        address sender;
        uint256 chainID;
    }

    /// @notice Struct containing revert context passed to onRevert.
    /// @param sender Address of account that initiated smart contract call.
    /// @param asset Address of asset, empty if it's gas token.
    /// @param amount Amount specified with the transaction.
    /// @param revertMessage Arbitrary data sent back in onRevert.
    struct RevertContext {
        address sender;
        address asset;
        uint256 amount;
        bytes revertMessage;
    }

    /// @notice Message context passed to execute function.
    /// @param sender Sender from omnichain contract.
    struct MessageContext {
        address sender;
    }

    // these structures allow to assess contract calls
    mapping(bytes32 => bool) public calledWithMessage;
    mapping(bytes => address) public senderWithMessage;
    mapping(bytes32 => uint256) public amountWithMessage;

    // the constructor is used to determine if the chain is ZetaChain
    constructor(bool isZetaChain_, address gateway_) {
        isZetaChain = isZetaChain_;
        gateway = gateway_;
    }

    // return the index used for the "WithMessage" mapping when the message for calls is empty
    // this allows testing the message with empty message
    // this function includes the sender of the message to avoid collisions when running parallel tests with different senders
    function getNoMessageIndex(address sender) pure public returns (string memory) {
        return string(abi.encodePacked(NO_MESSAGE_CALL, sender));
    }

    function setCalledWithMessage(string memory message) internal {
        calledWithMessage[keccak256(abi.encodePacked(message))] = true;
    }

    function setAmountWithMessage(string memory message, uint256 amount) internal {
        amountWithMessage[keccak256(abi.encodePacked(message))] = amount;
    }

    function getCalledWithMessage(string memory message) public view returns (bool) {
        return calledWithMessage[keccak256(abi.encodePacked(message))];
    }

    function getAmountWithMessage(string memory message) public view returns (uint256) {
        return amountWithMessage[keccak256(abi.encodePacked(message))];
    }

    // Universal contract interface on ZEVM
    function onCall(
        zContext calldata _context,
        address _zrc20,
        uint256 amount,
        bytes calldata message
    )
    external
    {
        require(!isRevertMessage(string(message)));

        string memory messageStr = message.length == 0 ? getNoMessageIndex(_context.sender) : string(message);

        setCalledWithMessage(messageStr);
        setAmountWithMessage(messageStr, amount);
    }

    // called with gas token
    function gasCall(string memory message) external payable {
        // Revert if the message is "revert"
        require(!isRevertMessage(message));

        setCalledWithMessage(message);
        setAmountWithMessage(message, msg.value);
    }

    // called with ERC20 token
    function erc20Call(IERC20 erc20, uint256 amount, string memory message) external {
        require(!isRevertMessage(message));
        require(erc20.transferFrom(msg.sender, address(this), amount));

        setCalledWithMessage(message);
        setAmountWithMessage(message, amount);
    }

    // called without token
    function simpleCall(string memory message) external {
        require(!isRevertMessage(message));

        setCalledWithMessage(message);
        setAmountWithMessage(message, 0);
    }

    // used to make functions revert
    function isRevertMessage(string memory message) internal pure returns (bool) {
        return keccak256(abi.encodePacked(message)) == keccak256(abi.encodePacked("revert"));
    }

    // Revertable interface
    function onRevert(RevertContext calldata revertContext) external {

        // if the chain is ZetaChain, consume gas to test the gas consumption
        // we do it specifically for ZetaChain to test the outbound processing workflow
        if (isZetaChain) {
            consumeGas();

            // withdraw funds to the sender on connected chain
            if (isWithdrawMessage(string(revertContext.revertMessage))) {
                (address feeToken, uint256 feeAmount) = IZRC20(revertContext.asset).withdrawGasFee();
                require(feeToken == revertContext.asset, "zrc20 is not gas token");
                require(feeAmount <= revertContext.amount, "fee amount is higher than the amount");
                uint256 withdrawAmount = revertContext.amount - feeAmount;

                IZRC20(revertContext.asset).approve(msg.sender, revertContext.amount);

                // caller is the gateway
                IGatewayZEVM(msg.sender).withdraw(
                    abi.encode(revertContext.sender),
                    withdrawAmount,
                    revertContext.asset,
                    RevertOptions(address(0), false, address(0), "", 0)
                );
            }
        }

        setCalledWithMessage(string(revertContext.revertMessage));
        setAmountWithMessage(string(revertContext.revertMessage), 0);
        senderWithMessage[revertContext.revertMessage] = revertContext.sender;
    }

    // Callable interface on connected EVM chains
    function onCall(MessageContext calldata messageContext, bytes calldata message) external payable returns (bytes memory) {
        string memory messageStr = message.length == 0 ? getNoMessageIndex(messageContext.sender) : string(message);

        setCalledWithMessage(messageStr);
        setAmountWithMessage(messageStr, msg.value);
        senderWithMessage[bytes(messageStr)] = messageContext.sender;

        return "";
    }

    // deposit through Gateway EVM
    function gatewayDeposit(address dst) external payable {
        require(!isZetaChain);
        IGatewayEVM(gateway).deposit{value: msg.value}(dst, RevertOptions(msg.sender, false, address(0), "", 0));
    }

    // deposit and call through Gateway EVM
    function gatewayDepositAndCall(address dst, bytes calldata payload) external payable {
        require(!isZetaChain);
        IGatewayEVM(gateway).depositAndCall{value: msg.value}(dst, payload, RevertOptions(msg.sender, false, address(0), "", 0));
    }

    // call through Gateway EVM
    function gatewayCall(address dst, bytes calldata payload) external {
        require(!isZetaChain);
        IGatewayEVM(gateway).call(dst, payload, RevertOptions(msg.sender, false, address(0), "", 0));
    }

    function consumeGas() internal {
        // Approximate target gas consumption
        uint256 targetGas = 500000;
        // Approximate gas cost for a single storage write
        uint256 storageWriteGasCost = 20000;
        uint256 iterations = targetGas / storageWriteGasCost;

        // Perform the storage writes
        for (uint256 i = 0; i < iterations; i++) {
            storageArray.push(i);
        }

        // Reset the storage array to avoid accumulation of storage cost
        delete storageArray;
    }

    function isWithdrawMessage(string memory message) internal pure returns (bool) {
        return keccak256(abi.encodePacked(message)) == keccak256(abi.encodePacked("withdraw"));
    }

    receive() external payable {}
}