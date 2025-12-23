package data

import (
	"github.com/multiversx/mx-chain-core-go/data"

	evmCore "github.com/multiversx/mx-chain-core-evm-go/core"
)

// EVMTransactionHandler defines the type of evm transaction
type EVMTransactionHandler interface {
	data.TransactionHandler
	GetRcvAliasAddr() []byte
	GetSndAliasAddr() []byte
	GetOriginalData() []byte

	GetMainAddressIdentifier() evmCore.AddressIdentifier
}
