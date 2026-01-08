package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/x402stacks/stacks-facilitator/internal/payment/domain/valueobject"
)

func createTestTransaction() BlockchainTransaction {
	txID, _ := valueobject.NewTransactionID("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	sender, _ := valueobject.NewStacksAddress("ST2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7")
	recipient, _ := valueobject.NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")

	return BlockchainTransaction{
		TxID:        txID,
		TokenType:   valueobject.TokenSTX,
		Sender:      sender,
		Recipient:   recipient,
		Amount:      valueobject.NewAmount(1000000),
		Fee:         valueobject.NewAmount(180),
		Nonce:       5,
		BlockHeight: 12345,
		Memo:        "test payment",
		Status:      "success",
		IsConfirmed: true,
	}
}

func TestVerificationService_ValidTransaction(t *testing.T) {
	svc := NewVerificationService()
	tx := createTestTransaction()
	recipient, _ := valueobject.NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")

	criteria := VerificationCriteria{
		ExpectedRecipient: recipient,
		MinAmount:         valueobject.NewAmount(500000),
		AcceptUnconfirmed: false,
	}

	result := svc.Verify(tx, criteria)

	assert.True(t, result.Valid)
	assert.Empty(t, result.Errors)
}

func TestVerificationService_RejectsWrongRecipient(t *testing.T) {
	svc := NewVerificationService()
	tx := createTestTransaction()
	wrongRecipient, _ := valueobject.NewStacksAddress("ST3J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ8")

	criteria := VerificationCriteria{
		ExpectedRecipient: wrongRecipient,
		MinAmount:         valueobject.NewAmount(500000),
		AcceptUnconfirmed: false,
	}

	result := svc.Verify(tx, criteria)

	assert.False(t, result.Valid)
	assert.Contains(t, result.Errors[0], "recipient mismatch")
}

func TestVerificationService_RejectsInsufficientAmount(t *testing.T) {
	svc := NewVerificationService()
	tx := createTestTransaction()
	recipient, _ := valueobject.NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")

	criteria := VerificationCriteria{
		ExpectedRecipient: recipient,
		MinAmount:         valueobject.NewAmount(2000000), // More than tx amount
		AcceptUnconfirmed: false,
	}

	result := svc.Verify(tx, criteria)

	assert.False(t, result.Valid)
	assert.Contains(t, result.Errors[0], "insufficient amount")
}

func TestVerificationService_RejectsUnconfirmedWhenRequired(t *testing.T) {
	svc := NewVerificationService()
	tx := createTestTransaction()
	tx.IsConfirmed = false
	tx.BlockHeight = 0
	recipient, _ := valueobject.NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")

	criteria := VerificationCriteria{
		ExpectedRecipient: recipient,
		MinAmount:         valueobject.NewAmount(500000),
		AcceptUnconfirmed: false,
	}

	result := svc.Verify(tx, criteria)

	assert.False(t, result.Valid)
	assert.Contains(t, result.Errors[0], "transaction not confirmed")
}

func TestVerificationService_AcceptsUnconfirmedWhenAllowed(t *testing.T) {
	svc := NewVerificationService()
	tx := createTestTransaction()
	tx.IsConfirmed = false
	tx.BlockHeight = 0
	recipient, _ := valueobject.NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")

	criteria := VerificationCriteria{
		ExpectedRecipient: recipient,
		MinAmount:         valueobject.NewAmount(500000),
		AcceptUnconfirmed: true,
	}

	result := svc.Verify(tx, criteria)

	assert.True(t, result.Valid)
}

func TestVerificationService_RejectsFailedTransaction(t *testing.T) {
	svc := NewVerificationService()
	tx := createTestTransaction()
	tx.Status = "failed"
	recipient, _ := valueobject.NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")

	criteria := VerificationCriteria{
		ExpectedRecipient: recipient,
		MinAmount:         valueobject.NewAmount(500000),
		AcceptUnconfirmed: false,
	}

	result := svc.Verify(tx, criteria)

	assert.False(t, result.Valid)
	assert.Contains(t, result.Errors[0], "transaction failed")
}

func TestVerificationService_RejectsAbortedTransaction(t *testing.T) {
	svc := NewVerificationService()
	tx := createTestTransaction()
	tx.Status = "abort_by_response"
	recipient, _ := valueobject.NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")

	criteria := VerificationCriteria{
		ExpectedRecipient: recipient,
		MinAmount:         valueobject.NewAmount(500000),
		AcceptUnconfirmed: false,
	}

	result := svc.Verify(tx, criteria)

	assert.False(t, result.Valid)
	assert.Contains(t, result.Errors[0], "transaction failed")
}

func TestVerificationService_ValidatesOptionalSender(t *testing.T) {
	svc := NewVerificationService()
	tx := createTestTransaction()
	recipient, _ := valueobject.NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")
	wrongSender, _ := valueobject.NewStacksAddress("ST9J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ9")

	criteria := VerificationCriteria{
		ExpectedRecipient: recipient,
		MinAmount:         valueobject.NewAmount(500000),
		ExpectedSender:    &wrongSender,
		AcceptUnconfirmed: false,
	}

	result := svc.Verify(tx, criteria)

	assert.False(t, result.Valid)
	assert.Contains(t, result.Errors[0], "sender mismatch")
}

func TestVerificationService_ValidatesOptionalMemo(t *testing.T) {
	svc := NewVerificationService()
	tx := createTestTransaction()
	recipient, _ := valueobject.NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")
	expectedMemo := "wrong memo"

	criteria := VerificationCriteria{
		ExpectedRecipient: recipient,
		MinAmount:         valueobject.NewAmount(500000),
		ExpectedMemo:      &expectedMemo,
		AcceptUnconfirmed: false,
	}

	result := svc.Verify(tx, criteria)

	assert.False(t, result.Valid)
	assert.Contains(t, result.Errors[0], "memo mismatch")
}

func TestVerificationService_MatchingMemo(t *testing.T) {
	svc := NewVerificationService()
	tx := createTestTransaction()
	recipient, _ := valueobject.NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")
	expectedMemo := "test payment"

	criteria := VerificationCriteria{
		ExpectedRecipient: recipient,
		MinAmount:         valueobject.NewAmount(500000),
		ExpectedMemo:      &expectedMemo,
		AcceptUnconfirmed: false,
	}

	result := svc.Verify(tx, criteria)

	assert.True(t, result.Valid)
}

func TestVerificationService_CollectsMultipleErrors(t *testing.T) {
	svc := NewVerificationService()
	tx := createTestTransaction()
	tx.Status = "failed"
	wrongRecipient, _ := valueobject.NewStacksAddress("ST3J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ8")

	criteria := VerificationCriteria{
		ExpectedRecipient: wrongRecipient,
		MinAmount:         valueobject.NewAmount(5000000),
		AcceptUnconfirmed: false,
	}

	result := svc.Verify(tx, criteria)

	assert.False(t, result.Valid)
	require.GreaterOrEqual(t, len(result.Errors), 2)
}
