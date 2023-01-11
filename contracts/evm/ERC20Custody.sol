// SPDX-License-Identifier: MIT
// v1.0, 2023-01-10
pragma solidity ^0.8.7;

// This ERC20 interface comes from OpenZeppelin
// https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/IERC20.sol
interface IERC20 {
    function transfer(address to, uint256 amount) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom( address from, address to, uint256 amount) external returns (bool);
    function balanceOf(address account) external view returns (uint256);
}

contract ERC20Custody {
    error NotWhitelisted();
    error NotPaused();
    error InvalidSender();
    error InvalidTSSUpdater();
    error ZeroAddress();
    error IsPaused();

    bool public paused;

    // TSSAddress is the TSS address collectively possessed by Zeta blockchain validators. 
    // Threshold Signature Scheme (TSS) [GG20] is a multi-sig ECDSA/EdDSA protocol. 
    address public TSSAddress; 
    address public TSSAddressUpdater;

    // Mapping of whitelisted token => true/false
    mapping(address => bool) public whitelisted;
    
    event Paused(address sender);
    event Unpaused(address sender);
    event Whitelisted(address asset);
    event Unwhitelisted(address asset);
    event Deposited(bytes recipient, address asset, uint256 amount, bytes message);
    event Withdrawn(address recipient, address asset, uint256 amount);

    constructor(address _TSSAddress, address _TSSAddressUpdater) {       
        TSSAddress = _TSSAddress; 
        TSSAddressUpdater = _TSSAddressUpdater; 
        paused = false; 
    }

    // update the TSSAddress in case of Zeta blockchain validator nodes churn
    function updateTSSAddress(address _address) external {
        if (msg.sender != TSSAddressUpdater) {
            revert InvalidTSSUpdater();
        }
        if (_address == address(0)) {
            revert ZeroAddress();
        }
        TSSAddress = _address;
    }

    // Change the ownership of TSSAddressUpdater to the Zeta blockchain TSS nodes. 
    // Effectively, only Zeta blockchain validators collectively can update TSSAddress afterwards. 
    function renounceTSSAddressUpdater() external {
        if (msg.sender != TSSAddressUpdater) {
            revert InvalidTSSUpdater();
        }
        if (TSSAddress == address(0)) {
            revert ZeroAddress();
        }
        TSSAddressUpdater = TSSAddress;
    }

    function pause() external {
        if (paused) {
            revert IsPaused();
        }
        if (msg.sender != TSSAddressUpdater) {
            revert InvalidTSSUpdater();
        }
        if (TSSAddress == address(0)) {
            revert ZeroAddress();
        }
        paused = true;
        emit Paused(msg.sender);
    }

    function unpause() external {
        if (!paused) {
            revert NotPaused();
        }
        if (msg.sender != TSSAddressUpdater) {
            revert InvalidTSSUpdater();
        }
        paused = false;
        emit Unpaused(msg.sender);
    }

    function whitelist(address asset) external {
        if (msg.sender != TSSAddress) {
            revert InvalidSender();
        }
        whitelisted[asset] = true;
        emit Whitelisted(asset);
    }

    function unwhitelist(address asset) external {
        if (msg.sender != TSSAddress) {
            revert InvalidSender();
        }
        whitelisted[asset] = false;
        emit Unwhitelisted(asset);
    }

    function deposit(bytes calldata recipient, address asset, uint256 amount, bytes calldata message) external {
        if (paused) {
            revert IsPaused();
        }
        if (!whitelisted[asset]) {
            revert NotWhitelisted();
        }
        IERC20(asset).transferFrom(msg.sender, address(this), amount);
        emit Deposited(recipient, asset, amount, message);
    }

    function withdraw(address recipient, address asset, uint256 amount) external {
        if (paused) {
            revert IsPaused();
        }
        if (msg.sender != TSSAddress) {
            revert InvalidSender();
        }
        if (!whitelisted[asset]) {
            revert NotWhitelisted();
        }
        IERC20(asset).transfer(recipient, amount);
        emit Withdrawn(recipient, asset, amount);
    }
}