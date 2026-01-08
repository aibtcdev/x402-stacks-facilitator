package command

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/x402stacks/stacks-facilitator/internal/payment/domain/service"
	"github.com/x402stacks/stacks-facilitator/internal/payment/domain/valueobject"
)

// MockBroadcaster is a mock implementation for testing
type MockBroadcaster struct {
	BroadcastFn         func(ctx context.Context, signedTx string, network valueobject.Network) (valueobject.TransactionID, error)
	WaitForConfirmFn    func(ctx context.Context, txID valueobject.TransactionID, tokenType valueobject.TokenType, network valueobject.Network, maxRetries int, retryDelay time.Duration) (service.BlockchainTransaction, error)
}

func (m *MockBroadcaster) BroadcastTransaction(ctx context.Context, signedTx string, network valueobject.Network) (valueobject.TransactionID, error) {
	return m.BroadcastFn(ctx, signedTx, network)
}

func (m *MockBroadcaster) WaitForConfirmation(ctx context.Context, txID valueobject.TransactionID, tokenType valueobject.TokenType, network valueobject.Network, maxRetries int, retryDelay time.Duration) (service.BlockchainTransaction, error) {
	return m.WaitForConfirmFn(ctx, txID, tokenType, network, maxRetries, retryDelay)
}

func TestSettlePaymentHandler_Success(t *testing.T) {
	txID, _ := valueobject.NewTransactionID("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	sender, _ := valueobject.NewStacksAddress("ST2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7")
	recipient, _ := valueobject.NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")

	mockTx := service.BlockchainTransaction{
		TxID:        txID,
		TokenType:   valueobject.TokenSTX,
		Sender:      sender,
		Recipient:   recipient,
		Amount:      valueobject.NewAmount(1000000),
		Fee:         valueobject.NewAmount(180),
		Nonce:       5,
		BlockHeight: 12345,
		Status:      "success",
		IsConfirmed: true,
	}

	mockBroadcaster := &MockBroadcaster{
		BroadcastFn: func(ctx context.Context, signedTx string, network valueobject.Network) (valueobject.TransactionID, error) {
			return txID, nil
		},
		WaitForConfirmFn: func(ctx context.Context, txID valueobject.TransactionID, tokenType valueobject.TokenType, network valueobject.Network, maxRetries int, retryDelay time.Duration) (service.BlockchainTransaction, error) {
			return mockTx, nil
		},
	}

	verificationSvc := service.NewVerificationService()
	handler := NewSettlePaymentHandler(mockBroadcaster, verificationSvc)

	cmd := SettlePaymentCommand{
		SignedTransaction: "0x00000001deadbeef",
		TokenType:         "STX",
		ExpectedRecipient: "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM",
		MinAmount:         500000,
		Network:           "testnet",
	}

	result, err := handler.Handle(context.Background(), cmd)

	require.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, "confirmed", result.Status)
	assert.Equal(t, txID.String(), result.TxID)
}

func TestSettlePaymentHandler_BroadcastError(t *testing.T) {
	mockBroadcaster := &MockBroadcaster{
		BroadcastFn: func(ctx context.Context, signedTx string, network valueobject.Network) (valueobject.TransactionID, error) {
			return valueobject.TransactionID{}, errors.New("broadcast failed")
		},
	}

	verificationSvc := service.NewVerificationService()
	handler := NewSettlePaymentHandler(mockBroadcaster, verificationSvc)

	cmd := SettlePaymentCommand{
		SignedTransaction: "0x00000001deadbeef",
		TokenType:         "STX",
		ExpectedRecipient: "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM",
		MinAmount:         500000,
		Network:           "testnet",
	}

	_, err := handler.Handle(context.Background(), cmd)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "broadcast")
}

func TestSettlePaymentHandler_VerificationFailed(t *testing.T) {
	txID, _ := valueobject.NewTransactionID("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	sender, _ := valueobject.NewStacksAddress("ST2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7")
	wrongRecipient, _ := valueobject.NewStacksAddress("ST9J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ9")

	mockTx := service.BlockchainTransaction{
		TxID:        txID,
		TokenType:   valueobject.TokenSTX,
		Sender:      sender,
		Recipient:   wrongRecipient, // Wrong recipient
		Amount:      valueobject.NewAmount(1000000),
		Fee:         valueobject.NewAmount(180),
		BlockHeight: 12345,
		Status:      "success",
		IsConfirmed: true,
	}

	mockBroadcaster := &MockBroadcaster{
		BroadcastFn: func(ctx context.Context, signedTx string, network valueobject.Network) (valueobject.TransactionID, error) {
			return txID, nil
		},
		WaitForConfirmFn: func(ctx context.Context, txID valueobject.TransactionID, tokenType valueobject.TokenType, network valueobject.Network, maxRetries int, retryDelay time.Duration) (service.BlockchainTransaction, error) {
			return mockTx, nil
		},
	}

	verificationSvc := service.NewVerificationService()
	handler := NewSettlePaymentHandler(mockBroadcaster, verificationSvc)

	cmd := SettlePaymentCommand{
		SignedTransaction: "0x00000001deadbeef",
		TokenType:         "STX",
		ExpectedRecipient: "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM",
		MinAmount:         500000,
		Network:           "testnet",
	}

	result, err := handler.Handle(context.Background(), cmd)

	require.NoError(t, err)
	assert.False(t, result.Success)
	assert.NotEmpty(t, result.Errors)
}

func TestSettlePaymentHandler_InvalidNetwork(t *testing.T) {
	mockBroadcaster := &MockBroadcaster{}
	verificationSvc := service.NewVerificationService()
	handler := NewSettlePaymentHandler(mockBroadcaster, verificationSvc)

	cmd := SettlePaymentCommand{
		SignedTransaction: "0x00000001deadbeef",
		TokenType:         "STX",
		ExpectedRecipient: "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM",
		MinAmount:         500000,
		Network:           "invalid",
	}

	_, err := handler.Handle(context.Background(), cmd)

	assert.Error(t, err)
}

func TestSettlePaymentHandler_InvalidRecipient(t *testing.T) {
	mockBroadcaster := &MockBroadcaster{}
	verificationSvc := service.NewVerificationService()
	handler := NewSettlePaymentHandler(mockBroadcaster, verificationSvc)

	cmd := SettlePaymentCommand{
		SignedTransaction: "0x00000001deadbeef",
		TokenType:         "STX",
		ExpectedRecipient: "invalid",
		MinAmount:         500000,
		Network:           "testnet",
	}

	_, err := handler.Handle(context.Background(), cmd)

	assert.Error(t, err)
}
