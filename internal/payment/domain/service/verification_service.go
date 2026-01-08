package service

import (
	"fmt"

	"github.com/x402stacks/stacks-facilitator/internal/payment/domain/valueobject"
)

// BlockchainTransaction represents a transaction fetched from the blockchain
type BlockchainTransaction struct {
	TxID        valueobject.TransactionID
	TokenType   valueobject.TokenType
	Sender      valueobject.StacksAddress
	Recipient   valueobject.StacksAddress
	Amount      valueobject.Amount
	Fee         valueobject.Amount
	Nonce       uint64
	BlockHeight uint64
	Memo        string
	Status      string
	IsConfirmed bool
}

// VerificationCriteria defines the criteria for validating a transaction
type VerificationCriteria struct {
	ExpectedRecipient valueobject.StacksAddress
	MinAmount         valueobject.Amount
	ExpectedSender    *valueobject.StacksAddress
	ExpectedMemo      *string
	AcceptUnconfirmed bool
}

// VerificationResult contains the result of a verification
type VerificationResult struct {
	Valid  bool
	Errors []string
}

// VerificationService validates blockchain transactions against criteria
type VerificationService struct{}

// NewVerificationService creates a new VerificationService
func NewVerificationService() *VerificationService {
	return &VerificationService{}
}

// Verify validates a blockchain transaction against the given criteria
func (s *VerificationService) Verify(tx BlockchainTransaction, criteria VerificationCriteria) VerificationResult {
	var errors []string

	// Check if transaction failed
	if isFailedStatus(tx.Status) {
		errors = append(errors, fmt.Sprintf("transaction failed with status: %s", tx.Status))
	}

	// Check confirmation requirement
	if !criteria.AcceptUnconfirmed && !tx.IsConfirmed {
		errors = append(errors, "transaction not confirmed")
	}

	// Check recipient
	if !tx.Recipient.Equals(criteria.ExpectedRecipient) {
		errors = append(errors, fmt.Sprintf("recipient mismatch: expected %s, got %s",
			criteria.ExpectedRecipient.String(), tx.Recipient.String()))
	}

	// Check amount
	if !tx.Amount.IsGreaterThanOrEqual(criteria.MinAmount) {
		errors = append(errors, fmt.Sprintf("insufficient amount: expected at least %s, got %s",
			criteria.MinAmount.String(), tx.Amount.String()))
	}

	// Check optional sender
	if criteria.ExpectedSender != nil && !tx.Sender.Equals(*criteria.ExpectedSender) {
		errors = append(errors, fmt.Sprintf("sender mismatch: expected %s, got %s",
			criteria.ExpectedSender.String(), tx.Sender.String()))
	}

	// Check optional memo
	if criteria.ExpectedMemo != nil && tx.Memo != *criteria.ExpectedMemo {
		errors = append(errors, fmt.Sprintf("memo mismatch: expected %s, got %s",
			*criteria.ExpectedMemo, tx.Memo))
	}

	return VerificationResult{
		Valid:  len(errors) == 0,
		Errors: errors,
	}
}

// isFailedStatus checks if the transaction status indicates failure
func isFailedStatus(status string) bool {
	failedStatuses := []string{
		"failed",
		"abort_by_response",
		"abort_by_post_condition",
	}
	for _, s := range failedStatuses {
		if status == s {
			return true
		}
	}
	return false
}
