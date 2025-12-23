package api

import (
	"github.com/multiversx/mx-chain-core-evm-go/core"
)

// AliasAddressRequest represents an alias address request for a given mvx address
type AliasAddressRequest struct {
	MvxAddress          string                 `json:"mvxAddress"`
	RequestedIdentifier core.AddressIdentifier `json:"requestedIdentifier"`
}

// MvxAddressRequest represents a mvx address request for a given alias address
type MvxAddressRequest struct {
	AliasAddress    string                 `json:"aliasAddress"`
	AliasIdentifier core.AddressIdentifier `json:"aliasIdentifier"`
}
