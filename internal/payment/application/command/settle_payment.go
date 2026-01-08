package command

import (
	"context"
	"fmt"
	"time"

	"github.com/x402stacks/stacks-facilitator/internal/payment/domain/service"
	"github.com/x402stacks/stacks-facilitator/internal/payment/domain/valueobject"
)

// TransactionBroadcaster interface for broadcasting and confirming transactions
type TransactionBroadcaster interface {
	BroadcastTransaction(ctx context.Context, signedTx string, network valueobject.Network) (valueobject.TransactionID, error)
	WaitForConfirmation(ctx context.Context, txID valueobject.TransactionID, tokenType valueobject.TokenType, network valueobject.Network, maxRetries int, retryDelay time.Duration) (service.BlockchainTransaction, error)
}

// SettlePaymentCommand represents a request to settle a payment
type SettlePaymentCommand struct {
	SignedTransaction string
	TokenType         string
	ExpectedRecipient string
	MinAmount         uint64
	ExpectedSender    *string
	Network           string
}

// SettlePaymentResult represents the result of a settlement
type SettlePaymentResult struct {
	Success          bool
	TxID             string
	SenderAddress    string
	RecipientAddress string
	Amount           uint64
	Fee              uint64
	Status           string
	BlockHeight      uint64
	TokenType        string
	Network          string
	Errors           []string
}

// SettlePaymentHandler handles settle payment commands
type SettlePaymentHandler struct {
	broadcaster     TransactionBroadcaster
	verificationSvc *service.VerificationService
	maxRetries      int
	retryDelay      time.Duration
}

// NewSettlePaymentHandler creates a new SettlePaymentHandler
func NewSettlePaymentHandler(broadcaster TransactionBroadcaster, verificationSvc *service.VerificationService) *SettlePaymentHandler {
	return &SettlePaymentHandler{
		broadcaster:     broadcaster,
		verificationSvc: verificationSvc,
		maxRetries:      15,
		retryDelay:      2 * time.Second,
	}
}

// Handle processes the settle payment command
func (h *SettlePaymentHandler) Handle(ctx context.Context, cmd SettlePaymentCommand) (SettlePaymentResult, error) {
	// Parse and validate inputs
	tokenType, err := valueobject.NewTokenType(cmd.TokenType)
	if err != nil {
		tokenType = valueobject.TokenSTX // Default to STX
	}

	network, err := valueobject.NewNetwork(cmd.Network)
	if err != nil {
		return SettlePaymentResult{}, fmt.Errorf("invalid network: %w", err)
	}

	expectedRecipient, err := valueobject.NewStacksAddress(cmd.ExpectedRecipient)
	if err != nil {
		return SettlePaymentResult{}, fmt.Errorf("invalid expected recipient: %w", err)
	}

	// Broadcast the transaction
	txID, err := h.broadcaster.BroadcastTransaction(ctx, cmd.SignedTransaction, network)
	if err != nil {
		return SettlePaymentResult{}, fmt.Errorf("failed to broadcast transaction: %w", err)
	}

	// Wait for transaction to be confirmed
	tx, err := h.broadcaster.WaitForConfirmation(ctx, txID, tokenType, network, h.maxRetries, h.retryDelay)
	if err != nil {
		return SettlePaymentResult{}, fmt.Errorf("failed to confirm transaction: %w", err)
	}

	// Build verification criteria (always require confirmation for settlement)
	criteria := service.VerificationCriteria{
		ExpectedRecipient: expectedRecipient,
		MinAmount:         valueobject.NewAmount(cmd.MinAmount),
		AcceptUnconfirmed: false,
	}

	// Optional sender
	if cmd.ExpectedSender != nil {
		sender, err := valueobject.NewStacksAddress(*cmd.ExpectedSender)
		if err == nil {
			criteria.ExpectedSender = &sender
		}
	}

	// Verify transaction
	verificationResult := h.verificationSvc.Verify(tx, criteria)

	// Determine status
	status := determinePaymentStatus(tx)

	return SettlePaymentResult{
		Success:          verificationResult.Valid,
		TxID:             tx.TxID.String(),
		SenderAddress:    tx.Sender.String(),
		RecipientAddress: tx.Recipient.String(),
		Amount:           tx.Amount.Value(),
		Fee:              tx.Fee.Value(),
		Status:           status,
		BlockHeight:      tx.BlockHeight,
		TokenType:        tx.TokenType.String(),
		Network:          network.String(),
		Errors:           verificationResult.Errors,
	}, nil
}
