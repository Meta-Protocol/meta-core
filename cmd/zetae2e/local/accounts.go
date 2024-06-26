package local

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
)

var (
	// DeployerAddress is the address of the account for deploying networks
	DeployerAddress    = ethcommon.HexToAddress("0xE5C5367B8224807Ac2207d350E60e1b6F27a7ecC")
	DeployerPrivateKey = "d87baf7bf6dc560a252596678c12e41f7d1682837f05b29d411bc3f78ae2c263" // #nosec G101 - used for testing

	// UserERC20Address is the address of the account for testing ERC20
	UserERC20Address    = ethcommon.HexToAddress("0x6F57D5E7c6DBb75e59F1524a3dE38Fc389ec5Fd6")
	UserERC20PrivateKey = "fda3be1b1517bdf48615bdadacc1e6463d2865868dc8077d2cdcfa4709a16894" // #nosec G101 - used for testing

	// UserZetaTestAddress is the address of the account for testing Zeta transfers
	UserZetaTestAddress    = ethcommon.HexToAddress("0x5cC2fBb200A929B372e3016F1925DcF988E081fd")
	UserZetaTestPrivateKey = "729a6cdc5c925242e7df92fdeeb94dadbf2d0b9950d4db8f034ab27a3b114ba7" // #nosec G101 - used for testing

	// UserZEVMMPTestAddress is the address of the account for testing ZEVM Message Passing
	UserZEVMMPTestAddress    = ethcommon.HexToAddress("0x8Ae229198eCE3c889C07DB648Ec7C30E6051592c")
	UserZEVMMPTestPrivateKey = "105460aebf71b10bfdb710ef5aa6d2932ee6ff6fc317ac9c24e0979903b10a5d" // #nosec G101 - used for testing

	// UserBitcoinAddress is the address of the account for testing Bitcoin
	UserBitcoinAddress    = ethcommon.HexToAddress("0x283d810090EdF4043E75247eAeBcE848806237fD")
	UserBitcoinPrivateKey = "7bb523963ee2c78570fb6113d886a4184d42565e8847f1cb639f5f5e2ef5b37a" // #nosec G101 - used for testing

	// UserEtherAddress is the address of the account for testing Ether
	UserEtherAddress    = ethcommon.HexToAddress("0x8D47Db7390AC4D3D449Cc20D799ce4748F97619A")
	UserEtherPrivateKey = "098e74a1c2261fa3c1b8cfca8ef2b4ff96c73ce36710d208d1f6535aef42545d" // #nosec G101 - used for testing

	// UserMiscAddress is the address of the account for miscellaneous tests
	UserMiscAddress    = ethcommon.HexToAddress("0x90126d02E41c9eB2a10cfc43aAb3BD3460523Cdf")
	UserMiscPrivateKey = "853c0945b8035a501b1161df65a17a0a20fc848bda8975a8b4e9222cc6f84cd4" // #nosec G101 - used for testing

	// UserAdminAddress is the address of the account for testing admin function features
	// NOTE: this is the default account using Anvil
	UserAdminAddress    = ethcommon.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	UserAdminPrivateKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" // #nosec G101 - used for testing

	// FungibleAdminAddress is the address of the account for testing the fungible admin functions
	UserFungibleAdminAddress    = ethcommon.HexToAddress("0x8305C114Ea73cAc4A88f39A173803F94741b9055")
	UserFungibleAdminPrivateKey = "d88d09a7d6849c15a36eb6931f9dd616091a63e9849a2cc86f309ba11fb8fec5" // #nosec G101 - used for testing
)
