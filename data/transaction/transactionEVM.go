//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=$GOPATH/src transactionEVM.proto
package transaction

import (
	"math/big"
	"strconv"

	ethCommon "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/transaction"

	evmCore "github.com/multiversx/mx-chain-core-evm-go/core"
)

// IsInterfaceNil verifies if underlying object is nil
func (tx *TransactionEVM) IsInterfaceNil() bool {
	return tx == nil
}

// GetValue returns the value of the transaction
func (tx *TransactionEVM) GetValue() *big.Int {
	if tx == nil {
		return nil
	}

	return tx.Transaction.GetValue()
}

// GetNonce returns the nonce of the transaction
func (tx *TransactionEVM) GetNonce() uint64 {
	if tx == nil {
		return 0
	}

	return tx.Transaction.GetNonce()
}

// GetData returns the data of the transaction
func (tx *TransactionEVM) GetData() []byte {
	if tx == nil {
		return nil
	}

	return tx.Transaction.GetData()
}

// GetRcvAddr returns the recevier address of the transaction
func (tx *TransactionEVM) GetRcvAddr() []byte {
	if tx == nil {
		return nil
	}

	return tx.Transaction.GetRcvAddr()
}

// GetRcvUserName returns the receiver username of the transaction
func (tx *TransactionEVM) GetRcvUserName() []byte {
	if tx == nil {
		return nil
	}

	return tx.Transaction.GetRcvUserName()
}

// GetSndAddr returns the sender address of the transaction
func (tx *TransactionEVM) GetSndAddr() []byte {
	if tx == nil {
		return nil
	}

	return tx.Transaction.GetSndAddr()
}

// GetGasLimit returns the gas limit of the transaction
func (tx *TransactionEVM) GetGasLimit() uint64 {
	if tx == nil {
		return 0
	}

	return tx.Transaction.GetGasLimit()
}

// GetGasPrice returns the gas price of the transaction
func (tx *TransactionEVM) GetGasPrice() uint64 {
	if tx == nil {
		return 0
	}

	return tx.Transaction.GetGasPrice()
}

// GetChainID returns the chain id of the transaction
func (tx *TransactionEVM) GetChainID() []byte {
	if tx == nil {
		return nil
	}

	return tx.Transaction.GetChainID()
}

// GetSndUserName returns the sender username of the transaction
func (tx *TransactionEVM) GetSndUserName() []byte {
	if tx == nil {
		return nil
	}

	return tx.Transaction.GetSndUserName()
}

// GetVersion returns the version of the transaction
func (tx *TransactionEVM) GetVersion() uint32 {
	if tx == nil {
		return 0
	}

	return tx.Transaction.GetVersion()
}

// GetSignature returns the signature of the transaction
func (tx *TransactionEVM) GetSignature() []byte {
	if tx == nil {
		return nil
	}

	return tx.Transaction.GetSignature()
}

// GetOptions returns the options of the transaction
func (tx *TransactionEVM) GetOptions() uint32 {
	if tx == nil {
		return 0
	}

	return tx.Transaction.GetOptions()
}

// GetGuardianAddr returns the guardian address of the transaction
func (tx *TransactionEVM) GetGuardianAddr() []byte {
	if tx == nil {
		return nil
	}

	return tx.Transaction.GetGuardianAddr()
}

// GetGuardianSignature returns the guardian signature of the transaction
func (tx *TransactionEVM) GetGuardianSignature() []byte {
	if tx == nil {
		return nil
	}

	return tx.Transaction.GetGuardianSignature()
}

// GetRelayerAddr returns the relayer address of the transaction
func (tx *TransactionEVM) GetRelayerAddr() []byte {
	if tx == nil {
		return nil
	}

	return tx.Transaction.GetRelayerAddr()
}

// GetRelayerSignature returns the relayer signature of the transaction
func (tx *TransactionEVM) GetRelayerSignature() []byte {
	if tx == nil {
		return nil
	}

	return tx.Transaction.GetRelayerSignature()
}

// SetValue sets the value of the transaction
func (tx *TransactionEVM) SetValue(value *big.Int) {
	tx.Transaction.SetValue(value)
}

// SetData sets the data of the transaction
func (tx *TransactionEVM) SetData(data []byte) {
	tx.Transaction.SetData(data)
}

// SetRcvAddr sets the receiver address of the transaction
func (tx *TransactionEVM) SetRcvAddr(addr []byte) {
	tx.Transaction.SetRcvAddr(addr)
}

// SetSndAddr sets the sender address of the transaction
func (tx *TransactionEVM) SetSndAddr(addr []byte) {
	tx.Transaction.SetSndAddr(addr)
}

func (tx *TransactionEVM) SetGasLimit(gasLimit uint64) {
	tx.Transaction.SetGasLimit(gasLimit)
}

func (tx *TransactionEVM) SetGasPrice(gasPrice uint64) {
	tx.Transaction.SetGasPrice(gasPrice)
}

func (tx *TransactionEVM) SetSignature(signature []byte) {
	tx.Transaction.SetSignature(signature)
}

func (tx *TransactionEVM) SetGuardianAddr(addr []byte) {
	tx.Transaction.SetGuardianAddr(addr)
}

func (tx *TransactionEVM) SetGuardianSignature(signature []byte) {
	tx.Transaction.SetGuardianSignature(signature)
}

// GetDataForSigning returns the serialized transaction having an empty signature field
func (tx *TransactionEVM) GetDataForSigning(encoder data.Encoder, marshaller data.Marshaller, hasher data.Hasher) ([]byte, error) {
	if tx.HasOptionETHTransactionFormat() {
		return tx.getETHDataForSigning()
	}

	if check.IfNil(encoder) {
		return nil, transaction.ErrNilEncoder
	}
	if check.IfNil(marshaller) {
		return nil, transaction.ErrNilMarshalizer
	}
	if check.IfNil(hasher) {
		return nil, transaction.ErrNilHasher
	}

	receiverAddr, err := encoder.Encode(tx.Transaction.GetRcvAddr())
	if err != nil {
		return nil, err
	}

	senderAddr, err := encoder.Encode(tx.Transaction.GetSndAddr())
	if err != nil {
		return nil, err
	}

	ftx := &transaction.FrontendTransaction{
		Nonce:            tx.Transaction.GetNonce(),
		Value:            tx.Transaction.GetValue().String(),
		Receiver:         receiverAddr,
		Sender:           senderAddr,
		GasPrice:         tx.Transaction.GetGasPrice(),
		GasLimit:         tx.Transaction.GetGasLimit(),
		SenderUsername:   tx.Transaction.GetSndAddr(),
		ReceiverUsername: tx.Transaction.GetRcvAddr(),
		Data:             tx.Transaction.GetData(),
		ChainID:          string(tx.Transaction.GetChainID()),
		Version:          tx.Transaction.GetVersion(),
		Options:          tx.Transaction.GetOptions(),
	}

	if len(tx.Transaction.GetGuardianAddr()) > 0 {
		guardianAddr, errGuardian := encoder.Encode(tx.Transaction.GetGuardianAddr())
		if errGuardian != nil {
			return nil, errGuardian
		}

		ftx.GuardianAddr = guardianAddr
	}

	if len(tx.Transaction.GetRelayerAddr()) > 0 {
		relayerAddr, errRelayer := encoder.Encode(tx.Transaction.GetRelayerAddr())
		if errRelayer != nil {
			return nil, errRelayer
		}

		ftx.RelayerAddr = relayerAddr
	}

	ftxBytes, err := marshaller.Marshal(ftx)
	if err != nil {
		return nil, err
	}

	shouldSignOnTxHash := tx.Transaction.GetVersion() > core.InitialVersionOfTransaction && tx.HasOptionHashSignSet()
	if !shouldSignOnTxHash {
		return ftxBytes, nil
	}

	ftxHash := hasher.Compute(string(ftxBytes))

	return ftxHash, nil
}

// HasOptionGuardianSet returns true if the guarded transaction option is set
func (tx *TransactionEVM) HasOptionGuardianSet() bool {
	return tx.Transaction.GetOptions()&transaction.MaskGuardedTransaction > 0
}

// HasOptionHashSignSet returns true if the signed with hash option is set
func (tx *TransactionEVM) HasOptionHashSignSet() bool {
	return tx.Transaction.GetOptions()&transaction.MaskSignedWithHash > 0
}

// HasOptionETHTransactionFormat returns true if the ETH transaction format option is set
func (tx *TransactionEVM) HasOptionETHTransactionFormat() bool {
	return tx.Transaction.GetOptions()&MaskETHTransactionFormat > 0
}

// CheckIntegrity checks for not nil fields and negative value
func (tx *TransactionEVM) CheckIntegrity() error {
	if tx.Transaction.GetSignature() == nil {
		return data.ErrNilSignature
	}
	if tx.Transaction.GetValue() == nil {
		return data.ErrNilValue
	}
	if tx.Transaction.GetValue().Sign() < 0 {
		return data.ErrNegativeValue
	}
	if len(tx.Transaction.GetRcvUserName()) > core.MaxUserNameLength {
		return data.ErrInvalidUserNameLength
	}
	if len(tx.Transaction.GetSndUserName()) > core.MaxUserNameLength {
		return data.ErrInvalidUserNameLength
	}

	return nil
}

// GetMainAddressIdentifier returns the main address identifier
func (tx *TransactionEVM) GetMainAddressIdentifier() evmCore.AddressIdentifier {
	if tx.HasOptionETHTransactionFormat() {
		return evmCore.ETHAddressIdentifier
	}
	return evmCore.MVXAddressIdentifier
}

func (tx *TransactionEVM) getETHDataForSigning() ([]byte, error) {
	signer, err := tx.BuildEthereumSigner()
	if err != nil {
		return nil, err
	}
	return signer.Hash(tx.BuildEthereumTransaction()).Bytes(), nil
}

func (tx *TransactionEVM) BuildEthereumSigner() (ethTypes.Signer, error) {
	chainId, err := strconv.Atoi(string(tx.Transaction.GetChainID()))
	if err != nil {
		return nil, err
	}
	return ethTypes.LatestSignerForChainID(big.NewInt(int64(chainId))), nil
}

func (tx *TransactionEVM) BuildEthereumTransaction() *ethTypes.Transaction {
	var to *ethCommon.Address
	if tx.RcvAliasAddr != nil {
		address := ethCommon.BytesToAddress(tx.RcvAliasAddr)
		to = &address
	}
	return ethTypes.NewTx(&ethTypes.LegacyTx{
		Nonce:    tx.Transaction.GetNonce(),
		GasPrice: new(big.Int).SetUint64(tx.Transaction.GetGasPrice()),
		Gas:      tx.Transaction.GasPrice,
		To:       to,
		Value:    tx.Transaction.GetValue(),
		Data:     tx.OriginalData,
	})
}
