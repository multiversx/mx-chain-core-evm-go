package core

import (
	"errors"
)

// ErrInvalidAddressIdentifier signals that an invalid address identifier has been provided
var ErrInvalidAddressIdentifier = errors.New("invalid address identifier")
