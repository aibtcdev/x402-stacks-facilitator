package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTransactionID_ValidID(t *testing.T) {
	validTxID := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

	txID, err := NewTransactionID(validTxID)

	require.NoError(t, err)
	assert.Equal(t, validTxID, txID.String())
}

func TestNewTransactionID_AddsPrefix(t *testing.T) {
	txIDWithoutPrefix := "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

	txID, err := NewTransactionID(txIDWithoutPrefix)

	require.NoError(t, err)
	assert.Equal(t, "0x"+txIDWithoutPrefix, txID.String())
}

func TestNewTransactionID_RejectsEmpty(t *testing.T) {
	_, err := NewTransactionID("")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "transaction ID cannot be empty")
}

func TestNewTransactionID_RejectsTooShort(t *testing.T) {
	_, err := NewTransactionID("0x1234")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid transaction ID length")
}

func TestNewTransactionID_RejectsInvalidHex(t *testing.T) {
	invalidHex := "0xGGGG567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

	_, err := NewTransactionID(invalidHex)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid hex")
}

func TestTransactionID_Equals(t *testing.T) {
	txID1, _ := NewTransactionID("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	txID2, _ := NewTransactionID("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	txID3, _ := NewTransactionID("0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890")

	assert.True(t, txID1.Equals(txID2))
	assert.False(t, txID1.Equals(txID3))
}
