package valueobject

import "fmt"

// Amount represents a monetary amount in base units (microSTX, satoshis, etc.)
type Amount struct {
	value uint64
}

// NewAmount creates a new Amount
func NewAmount(value uint64) Amount {
	return Amount{value: value}
}

// Value returns the raw uint64 value
func (a Amount) Value() uint64 {
	return a.value
}

// IsZero checks if the amount is zero
func (a Amount) IsZero() bool {
	return a.value == 0
}

// IsGreaterThanOrEqual checks if this amount is >= other
func (a Amount) IsGreaterThanOrEqual(other Amount) bool {
	return a.value >= other.value
}

// Add adds two amounts
func (a Amount) Add(other Amount) Amount {
	return Amount{value: a.value + other.value}
}

// Subtract subtracts other from this amount
func (a Amount) Subtract(other Amount) Amount {
	if other.value > a.value {
		return Amount{value: 0}
	}
	return Amount{value: a.value - other.value}
}

// ToSTX converts microSTX to STX (1 STX = 1,000,000 microSTX)
func (a Amount) ToSTX() float64 {
	return float64(a.value) / 1_000_000
}

// String returns the amount as a string
func (a Amount) String() string {
	return fmt.Sprintf("%d", a.value)
}
