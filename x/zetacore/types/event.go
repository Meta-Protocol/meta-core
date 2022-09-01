package types

const (
	// event key
	SubTypeKey = "SubTypeKey"
	CctxIndex  = "CctxIndex"

	Sender        = "Sender"
	SenderChain   = "SenderChain"
	InTxHash      = "InTxObservedHash"
	InBlockHeight = "InTxObservedBlockHeight"

	Receiver      = "Receiver"
	ReceiverChain = "ReceiverChain"
	OutTxHash     = "OutTxObservedHash"

	ZetaMint       = "ZetaMint"
	ZetaBurnt      = "ZetaBurnt"
	OutBoundChain  = "OutBoundChain"
	OldStatus      = "OldStatus"
	NewStatus      = "NewStatus"
	StatusMessage  = "StatusMessage"
	RelayedMessage = "RelayedMessage"
	Identifiers    = "LogIdentifiers"
)

const (
	OutboundTxSuccessful = "zetacore/OutboundTxSuccessful"
	OutboundTxFailed     = "zetacore/OutboundTxFailed"
	InboundCreated       = "zetacore/InboundCreated"
	InboundFinalized     = "zetacore/InboundFinalized"
	SendScrubbed         = "SendScrubbed"
)
