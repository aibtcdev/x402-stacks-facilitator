package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNetwork_Mainnet(t *testing.T) {
	network, err := NewNetwork("mainnet")

	require.NoError(t, err)
	assert.Equal(t, NetworkMainnet, network)
}

func TestNewNetwork_Testnet(t *testing.T) {
	network, err := NewNetwork("testnet")

	require.NoError(t, err)
	assert.Equal(t, NetworkTestnet, network)
}

func TestNewNetwork_CaseInsensitive(t *testing.T) {
	network, err := NewNetwork("MAINNET")

	require.NoError(t, err)
	assert.Equal(t, NetworkMainnet, network)
}

func TestNewNetwork_Invalid(t *testing.T) {
	_, err := NewNetwork("devnet")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported network")
}

func TestNewNetwork_Empty(t *testing.T) {
	_, err := NewNetwork("")

	assert.Error(t, err)
}

func TestNetwork_String(t *testing.T) {
	assert.Equal(t, "mainnet", NetworkMainnet.String())
	assert.Equal(t, "testnet", NetworkTestnet.String())
}

func TestNetwork_APIBaseURL(t *testing.T) {
	assert.Equal(t, "https://api.mainnet.hiro.so", NetworkMainnet.APIBaseURL())
	assert.Equal(t, "https://api.testnet.hiro.so", NetworkTestnet.APIBaseURL())
}

func TestNetwork_IsMainnet(t *testing.T) {
	assert.True(t, NetworkMainnet.IsMainnet())
	assert.False(t, NetworkTestnet.IsMainnet())
}

func TestNetwork_IsTestnet(t *testing.T) {
	assert.False(t, NetworkMainnet.IsTestnet())
	assert.True(t, NetworkTestnet.IsTestnet())
}
