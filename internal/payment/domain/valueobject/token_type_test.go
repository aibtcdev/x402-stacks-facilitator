package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTokenType_STX(t *testing.T) {
	tokenType, err := NewTokenType("STX")

	require.NoError(t, err)
	assert.Equal(t, TokenSTX, tokenType)
}

func TestNewTokenType_SBTC(t *testing.T) {
	tokenType, err := NewTokenType("SBTC")

	require.NoError(t, err)
	assert.Equal(t, TokenSBTC, tokenType)
}

func TestNewTokenType_USDCX(t *testing.T) {
	tokenType, err := NewTokenType("USDCX")

	require.NoError(t, err)
	assert.Equal(t, TokenUSDCX, tokenType)
}

func TestNewTokenType_CaseInsensitive(t *testing.T) {
	tokenType, err := NewTokenType("stx")

	require.NoError(t, err)
	assert.Equal(t, TokenSTX, tokenType)
}

func TestNewTokenType_Invalid(t *testing.T) {
	_, err := NewTokenType("BTC")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported token type")
}

func TestNewTokenType_Empty(t *testing.T) {
	_, err := NewTokenType("")

	assert.Error(t, err)
}

func TestTokenType_String(t *testing.T) {
	assert.Equal(t, "STX", TokenSTX.String())
	assert.Equal(t, "SBTC", TokenSBTC.String())
	assert.Equal(t, "USDCX", TokenUSDCX.String())
}

func TestTokenType_IsNative(t *testing.T) {
	assert.True(t, TokenSTX.IsNative())
	assert.False(t, TokenSBTC.IsNative())
	assert.False(t, TokenUSDCX.IsNative())
}
