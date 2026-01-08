package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPaymentStatus_Pending(t *testing.T) {
	status, err := NewPaymentStatus("pending")

	require.NoError(t, err)
	assert.Equal(t, StatusPending, status)
}

func TestNewPaymentStatus_Confirmed(t *testing.T) {
	status, err := NewPaymentStatus("confirmed")

	require.NoError(t, err)
	assert.Equal(t, StatusConfirmed, status)
}

func TestNewPaymentStatus_Failed(t *testing.T) {
	status, err := NewPaymentStatus("failed")

	require.NoError(t, err)
	assert.Equal(t, StatusFailed, status)
}

func TestNewPaymentStatus_Invalid(t *testing.T) {
	_, err := NewPaymentStatus("unknown")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid payment status")
}

func TestNewPaymentStatus_Empty(t *testing.T) {
	_, err := NewPaymentStatus("")

	assert.Error(t, err)
}

func TestPaymentStatus_String(t *testing.T) {
	assert.Equal(t, "pending", StatusPending.String())
	assert.Equal(t, "confirmed", StatusConfirmed.String())
	assert.Equal(t, "failed", StatusFailed.String())
}

func TestPaymentStatus_IsConfirmed(t *testing.T) {
	assert.False(t, StatusPending.IsConfirmed())
	assert.True(t, StatusConfirmed.IsConfirmed())
	assert.False(t, StatusFailed.IsConfirmed())
}

func TestPaymentStatus_IsFailed(t *testing.T) {
	assert.False(t, StatusPending.IsFailed())
	assert.False(t, StatusConfirmed.IsFailed())
	assert.True(t, StatusFailed.IsFailed())
}

func TestPaymentStatus_IsPending(t *testing.T) {
	assert.True(t, StatusPending.IsPending())
	assert.False(t, StatusConfirmed.IsPending())
	assert.False(t, StatusFailed.IsPending())
}
