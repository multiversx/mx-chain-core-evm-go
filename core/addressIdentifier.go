package core

import (
	"encoding/hex"
	"encoding/json"
)

const (
	addressIdentifierSize      = 2
	addressIdentifierSeparator = "@"
)

type AddressIdentifier [addressIdentifierSize]byte

var (
	InvalidAddressIdentifier = AddressIdentifier{}
	MVXAddressIdentifier     = AddressIdentifier{0, 1}
	ETHAddressIdentifier     = AddressIdentifier{0, 2}
)

func (a AddressIdentifier) BuildAddressIdentifier(address []byte) string {
	return a.String() + addressIdentifierSeparator + string(address)
}

func (a AddressIdentifier) String() string {
	return string(a.Spread())
}

func (a AddressIdentifier) Spread() []byte {
	return a[:]
}

func (a *AddressIdentifier) UnmarshalJSON(serializedAddressIdentifier []byte) error {
	var addressIdentifierHex string
	if err := json.Unmarshal(serializedAddressIdentifier, &addressIdentifierHex); err != nil {
		return err
	}

	addressIdentifier, err := hex.DecodeString(addressIdentifierHex)
	if err != nil {
		return err
	}

	parsedAddressIdentifier := ParseAddressIdentifier(addressIdentifier)
	if parsedAddressIdentifier == InvalidAddressIdentifier {
		return ErrInvalidAddressIdentifier
	}

	*a = parsedAddressIdentifier
	return nil
}

func (a *AddressIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(a.Spread()))
}

func ParseAddressIdentifier(identifier []byte) AddressIdentifier {
	if len(identifier) != addressIdentifierSize {
		return InvalidAddressIdentifier
	}

	switch parsedAddressIdentifier := AddressIdentifier(identifier); parsedAddressIdentifier {
	case MVXAddressIdentifier, ETHAddressIdentifier:
		return parsedAddressIdentifier
	default:
		return InvalidAddressIdentifier
	}
}
