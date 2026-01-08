package command

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/x402stacks/stacks-facilitator/internal/payment/domain/service"
	"github.com/x402stacks/stacks-facilitator/internal/payment/domain/valueobject"
)

// MockBlockchainClient is a mock implementation for testing
type MockBlockchainClient struct {
	GetTransactionFn func(ctx context.Context, txID valueobject.TransactionID, tokenType valueobject.TokenType, network valueobject.Network) (service.BlockchainTransaction, error)
}

func (m *MockBlockchainClient) GetTransactionWithRetry(ctx context.Context, txID valueobject.TransactionID, tokenType valueobject.TokenType, network valueobject.Network, maxRetries int, retryDelay time.Duration) (service.BlockchainTransaction, error) {
	return m.GetTransactionFn(ctx, txID, tokenType, network)
}

func createMockTransaction() service.BlockchainTransaction {
	txID, _ := valueobject.NewTransactionID("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	sender, _ := valueobject.NewStacksAddress("ST2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7")
	recipient, _ := valueobject.NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")

	return service.BlockchainTransaction{
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

func TestVerifyPaymentHandler_Success(t *testing.T) {
	mockTx := createMockTransaction()
	mockClient := &MockBlockchainClient{
		GetTransactionFn: func(ctx context.Context, txID valueobject.TransactionID, tokenType valueobject.TokenType, network valueobject.Network) (service.BlockchainTransaction, error) {
			return mockTx, nil
		},
	}

	verificationSvc := service.NewVerificationService()
	handler := NewVerifyPaymentHandler(mockClient, verificationSvc)

	cmd := VerifyPaymentCommand{
		TxID:              "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		TokenType:         "STX",
		ExpectedRecipient: "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM",
		MinAmount:         500000,
		Network:           "testnet",
	}

	result, err := handler.Handle(context.Background(), cmd)

	require.NoError(t, err)
	assert.True(t, result.Valid)
	assert.Equal(t, "confirmed", result.Status)
	assert.Equal(t, uint64(1000000), result.Amount)
	assert.Equal(t, "ST2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7", result.SenderAddress)
}

func TestVerifyPaymentHandler_InvalidRecipient(t *testing.T) {
	mockTx := createMockTransaction()
	mockClient := &MockBlockchainClient{
		GetTransactionFn: func(ctx context.Context, txID valueobject.TransactionID, tokenType valueobject.TokenType, network valueobject.Network) (service.BlockchainTransaction, error) {
			return mockTx, nil
		},
	}

	verificationSvc := service.NewVerificationService()
	handler := NewVerifyPaymentHandler(mockClient, verificationSvc)

	cmd := VerifyPaymentCommand{
		TxID:              "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		TokenType:         "STX",
		ExpectedRecipient: "ST9J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ9", // Wrong recipient
		MinAmount:         500000,
		Network:           "testnet",
	}

	result, err := handler.Handle(context.Background(), cmd)

	require.NoError(t, err)
	assert.False(t, result.Valid)
	assert.NotEmpty(t, result.Errors)
}

func TestVerifyPaymentHandler_InsufficientAmount(t *testing.T) {
	mockTx := createMockTransaction()
	mockClient := &MockBlockchainClient{
		GetTransactionFn: func(ctx context.Context, txID valueobject.TransactionID, tokenType valueobject.TokenType, network valueobject.Network) (service.BlockchainTransaction, error) {
			return mockTx, nil
		},
	}

	verificationSvc := service.NewVerificationService()
	handler := NewVerifyPaymentHandler(mockClient, verificationSvc)

	cmd := VerifyPaymentCommand{
		TxID:              "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		TokenType:         "STX",
		ExpectedRecipient: "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM",
		MinAmount:         5000000, // More than tx amount
		Network:           "testnet",
	}

	result, err := handler.Handle(context.Background(), cmd)

	require.NoError(t, err)
	assert.False(t, result.Valid)
}

func TestVerifyPaymentHandler_WithOptionalSender(t *testing.T) {
	mockTx := createMockTransaction()
	mockClient := &MockBlockchainClient{
		GetTransactionFn: func(ctx context.Context, txID valueobject.TransactionID, tokenType valueobject.TokenType, network valueobject.Network) (service.BlockchainTransaction, error) {
			return mockTx, nil
		},
	}

	verificationSvc := service.NewVerificationService()
	handler := NewVerifyPaymentHandler(mockClient, verificationSvc)

	expectedSender := "ST2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7"
	cmd := VerifyPaymentCommand{
		TxID:              "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		TokenType:         "STX",
		ExpectedRecipient: "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM",
		MinAmount:         500000,
		ExpectedSender:    &expectedSender,
		Network:           "testnet",
	}

	result, err := handler.Handle(context.Background(), cmd)

	require.NoError(t, err)
	assert.True(t, result.Valid)
}

func TestVerifyPaymentHandler_WithOptionalMemo(t *testing.T) {
	mockTx := createMockTransaction()
	mockClient := &MockBlockchainClient{
		GetTransactionFn: func(ctx context.Context, txID valueobject.TransactionID, tokenType valueobject.TokenType, network valueobject.Network) (service.BlockchainTransaction, error) {
			return mockTx, nil
		},
	}

	verificationSvc := service.NewVerificationService()
	handler := NewVerifyPaymentHandler(mockClient, verificationSvc)

	expectedMemo := "test payment"
	cmd := VerifyPaymentCommand{
		TxID:              "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		TokenType:         "STX",
		ExpectedRecipient: "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM",
		MinAmount:         500000,
		ExpectedMemo:      &expectedMemo,
		Network:           "testnet",
	}

	result, err := handler.Handle(context.Background(), cmd)

	require.NoError(t, err)
	assert.True(t, result.Valid)
}

func TestVerifyPaymentHandler_InvalidTxID(t *testing.T) {
	mockClient := &MockBlockchainClient{}
	verificationSvc := service.NewVerificationService()
	handler := NewVerifyPaymentHandler(mockClient, verificationSvc)

	cmd := VerifyPaymentCommand{
		TxID:              "invalid",
		TokenType:         "STX",
		ExpectedRecipient: "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM",
		MinAmount:         500000,
		Network:           "testnet",
	}

	_, err := handler.Handle(context.Background(), cmd)

	assert.Error(t, err)
}

func TestVerifyPaymentHandler_InvalidNetwork(t *testing.T) {
	mockClient := &MockBlockchainClient{}
	verificationSvc := service.NewVerificationService()
	handler := NewVerifyPaymentHandler(mockClient, verificationSvc)

	cmd := VerifyPaymentCommand{
		TxID:              "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		TokenType:         "STX",
		ExpectedRecipient: "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM",
		MinAmount:         500000,
		Network:           "invalid",
	}

	_, err := handler.Handle(context.Background(), cmd)

	assert.Error(t, err)
}
