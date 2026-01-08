package valueobject

import (
	"encoding/hex"
	"errors"
	"strings"
)

// TransactionID represents a Stacks blockchain transaction ID
type TransactionID struct {
	value string
}

// NewTransactionID creates a new TransactionID from a string
func NewTransactionID(id string) (TransactionID, error) {
	if id == "" {
		return TransactionID{}, errors.New("transaction ID cannot be empty")
	}

	// Normalize: add 0x prefix if missing
	normalized := id
	if !strings.HasPrefix(id, "0x") {
		normalized = "0x" + id
	}

	// Validate length (0x + 64 hex chars = 66)
	if len(normalized) != 66 {
		return TransactionID{}, errors.New("invalid transaction ID length: expected 66 characters")
	}

	// Validate hex characters
	hexPart := normalized[2:]
	if _, err := hex.DecodeString(hexPart); err != nil {
		return TransactionID{}, errors.New("invalid hex characters in transaction ID")
	}

	return TransactionID{value: normalized}, nil
}

// String returns the transaction ID as a string
func (t TransactionID) String() string {
	return t.value
}

// Equals checks if two TransactionIDs are equal
func (t TransactionID) Equals(other TransactionID) bool {
	return t.value == other.value
}

// IsZero checks if the TransactionID is empty
func (t TransactionID) IsZero() bool {
	return t.value == ""
}
