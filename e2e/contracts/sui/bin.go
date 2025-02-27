package sui

import (
	_ "embed"
	"encoding/base64"
)

//go:embed gateway.mv
var gatewayBinary []byte

//go:embed fake_usdc.mv
var fakeUSDC []byte

// GatewayBytecodeBase64 gets the gateway binary encoded as base64 for deployment
func GatewayBytecodeBase64() string {
	return base64.StdEncoding.EncodeToString(gatewayBinary)
}

// FakeUSDCBytecodeBase64 gets the fake USDC binary encoded as base64 for deployment
func FakeUSDCBytecodeBase64() string {
	return base64.StdEncoding.EncodeToString(fakeUSDC)
}
