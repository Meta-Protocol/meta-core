// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;
interface ZetaInterfaces {
    /**
     * @dev Use SendInput to interact with the Connector: connector.send(SendInput)
     */
    struct SendInput {
        /// @dev Chain id of the destination chain. More about chain ids https://docs.zetachain.com/learn/glossary#chain-id
        uint256 destinationChainId;
        /// @dev Address receiving the message on the destination chain (expressed in bytes since it can be non-EVM)
        bytes destinationAddress;
        /// @dev Gas limit for the destination chain's transaction
        uint256 destinationGasLimit;
        /// @dev An encoded, arbitrary message to be parsed by the destination contract
        bytes message;
        /// @dev ZETA to be sent cross-chain + ZetaChain gas fees + destination chain gas fees (expressed in ZETA)
        uint256 zetaValueAndGas;
        /// @dev Optional parameters for the ZetaChain protocol
        bytes zetaParams;
    }

    /**
     * @dev Our Connector calls onZetaMessage with this struct as argument
     */
    struct ZetaMessage {
        bytes zetaTxSenderAddress;
        uint256 sourceChainId;
        address destinationAddress;
        /// @dev Remaining ZETA from zetaValueAndGas after subtracting ZetaChain gas fees and destination gas fees
        uint256 zetaValue;
        bytes message;
    }

    /**
     * @dev Our Connector calls onZetaRevert with this struct as argument
     */
    struct ZetaRevert {
        address zetaTxSenderAddress;
        uint256 sourceChainId;
        bytes destinationAddress;
        uint256 destinationChainId;
        /// @dev Equals to: zetaValueAndGas - ZetaChain gas fees - destination chain gas fees - source chain revert tx gas fees
        uint256 remainingZetaValue;
        bytes message;
    }
}

interface ZetaConnector {
    /**
     * @dev Sending value and data cross-chain is as easy as calling connector.send(SendInput)
     */
    function send(ZetaInterfaces.SendInput calldata input) external;
}

interface ZetaReceiver {
    /**
     * @dev onZetaMessage is called when a cross-chain message reaches a contract
     */
    function onZetaMessage(ZetaInterfaces.ZetaMessage calldata zetaMessage) external;

    /**
     * @dev onZetaRevert is called when a cross-chain message reverts.
     * It's useful to rollback to the original state
     */
    function onZetaRevert(ZetaInterfaces.ZetaRevert calldata zetaRevert) external;
}

contract TestDApp is  ZetaReceiver {
    bytes32 public constant HELLO_WORLD_MESSAGE_TYPE = keccak256("CROSS_CHAIN_HELLO_WORLD");
    event HelloWorldEvent();
    event RevertedHelloWorldEvent();
    error InvalidMessageType();
    address public connector;
    constructor(address _connector) {
        connector = _connector;
    }

    function onZetaMessage(ZetaInterfaces.ZetaMessage calldata zetaMessage) external override {
        (, bool doRevert) = abi.decode(zetaMessage.message, (bytes32, bool));
        require(doRevert == false,  "message says revert");

        emit HelloWorldEvent();
    }

    function sendHelloWorld(address destinationAddress, uint256 destinationChainId, uint256 value, bool doRevert) external payable {
        ZetaConnector(connector).send(
            ZetaInterfaces.SendInput({
                destinationChainId: destinationChainId,
                destinationAddress: abi.encodePacked(destinationAddress),
                destinationGasLimit: 250000,
                message: abi.encode(HELLO_WORLD_MESSAGE_TYPE, doRevert),
                zetaValueAndGas: value,
                zetaParams: abi.encode("")
            })
        );
    }

    function onZetaRevert(ZetaInterfaces.ZetaRevert calldata zetaRevert) external override {
        (, bool doRevert) = abi.decode(zetaRevert.message, (bytes32, bool));
        require(doRevert == true, "the 1st outbound was not caused by revert flag in message");
        emit RevertedHelloWorldEvent();
    }
}