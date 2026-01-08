package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAmount_Valid(t *testing.T) {
	amount := NewAmount(1000000)

	assert.Equal(t, uint64(1000000), amount.Value())
}

func TestNewAmount_Zero(t *testing.T) {
	amount := NewAmount(0)

	assert.Equal(t, uint64(0), amount.Value())
	assert.True(t, amount.IsZero())
}

func TestAmount_IsGreaterThanOrEqual(t *testing.T) {
	amount1 := NewAmount(1000000)
	amount2 := NewAmount(500000)
	amount3 := NewAmount(1000000)

	assert.True(t, amount1.IsGreaterThanOrEqual(amount2))
	assert.True(t, amount1.IsGreaterThanOrEqual(amount3))
	assert.False(t, amount2.IsGreaterThanOrEqual(amount1))
}

func TestAmount_Add(t *testing.T) {
	amount1 := NewAmount(1000000)
	amount2 := NewAmount(500000)

	result := amount1.Add(amount2)

	assert.Equal(t, uint64(1500000), result.Value())
}

func TestAmount_Subtract(t *testing.T) {
	amount1 := NewAmount(1000000)
	amount2 := NewAmount(500000)

	result := amount1.Subtract(amount2)

	assert.Equal(t, uint64(500000), result.Value())
}

func TestAmount_ToSTX(t *testing.T) {
	// 1 STX = 1,000,000 microSTX
	amount := NewAmount(1500000)

	stx := amount.ToSTX()

	assert.Equal(t, 1.5, stx)
}

func TestAmount_String(t *testing.T) {
	amount := NewAmount(1000000)

	assert.Equal(t, "1000000", amount.String())
}
