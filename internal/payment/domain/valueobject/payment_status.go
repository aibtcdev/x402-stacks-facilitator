package valueobject

import (
	"errors"
	"strings"
)

// PaymentStatus represents the status of a payment
type PaymentStatus string

const (
	StatusPending   PaymentStatus = "pending"
	StatusConfirmed PaymentStatus = "confirmed"
	StatusFailed    PaymentStatus = "failed"
)

// NewPaymentStatus creates a new PaymentStatus from a string
func NewPaymentStatus(s string) (PaymentStatus, error) {
	if s == "" {
		return "", errors.New("payment status cannot be empty")
	}

	normalized := strings.ToLower(s)
	switch normalized {
	case "pending":
		return StatusPending, nil
	case "confirmed":
		return StatusConfirmed, nil
	case "failed":
		return StatusFailed, nil
	default:
		return "", errors.New("invalid payment status: " + s)
	}
}

// String returns the status as a string
func (s PaymentStatus) String() string {
	return string(s)
}

// IsConfirmed returns true if status is confirmed
func (s PaymentStatus) IsConfirmed() bool {
	return s == StatusConfirmed
}

// IsFailed returns true if status is failed
func (s PaymentStatus) IsFailed() bool {
	return s == StatusFailed
}

// IsPending returns true if status is pending
func (s PaymentStatus) IsPending() bool {
	return s == StatusPending
}
