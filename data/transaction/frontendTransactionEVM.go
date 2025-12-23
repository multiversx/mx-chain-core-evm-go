package transaction

import (
	"github.com/multiversx/mx-chain-core-go/data/transaction"
)

type FrontendTransactionEVM struct {
	*transaction.FrontendTransaction
	SenderAliasAddress   []byte `json:"senderAliasAddress,omitempty"`
	ReceiverAliasAddress []byte `json:"receiverAliasAddress,omitempty"`
	OriginalData         []byte `json:"originalData,omitempty"`
}
