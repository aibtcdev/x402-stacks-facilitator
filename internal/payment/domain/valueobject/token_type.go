package valueobject

import (
	"errors"
	"strings"
)

// TokenType represents a supported token type
type TokenType string

const (
	TokenSTX   TokenType = "STX"
	TokenSBTC  TokenType = "SBTC"
	TokenUSDCX TokenType = "USDCX"
)

// NewTokenType creates a new TokenType from a string
func NewTokenType(s string) (TokenType, error) {
	if s == "" {
		return "", errors.New("token type cannot be empty")
	}

	normalized := strings.ToUpper(s)
	switch normalized {
	case "STX":
		return TokenSTX, nil
	case "SBTC":
		return TokenSBTC, nil
	case "USDCX":
		return TokenUSDCX, nil
	default:
		return "", errors.New("unsupported token type: " + s)
	}
}

// String returns the token type as a string
func (t TokenType) String() string {
	return string(t)
}

// IsNative returns true if this is the native STX token
func (t TokenType) IsNative() bool {
	return t == TokenSTX
}

// IsSIP010 returns true if this is a SIP-010 token (not native STX)
func (t TokenType) IsSIP010() bool {
	return t == TokenSBTC || t == TokenUSDCX
}
