package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStacksAddress_ValidTestnetAddress(t *testing.T) {
	validAddr := "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM"

	addr, err := NewStacksAddress(validAddr)

	require.NoError(t, err)
	assert.Equal(t, validAddr, addr.String())
}

func TestNewStacksAddress_ValidMainnetAddress(t *testing.T) {
	validAddr := "SP2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKQ9H6DPR"

	addr, err := NewStacksAddress(validAddr)

	require.NoError(t, err)
	assert.Equal(t, validAddr, addr.String())
}

func TestNewStacksAddress_RejectsEmpty(t *testing.T) {
	_, err := NewStacksAddress("")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "address cannot be empty")
}

func TestNewStacksAddress_RejectsInvalidPrefix(t *testing.T) {
	_, err := NewStacksAddress("BT1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid Stacks address prefix")
}

func TestNewStacksAddress_RejectsTooShort(t *testing.T) {
	_, err := NewStacksAddress("ST1234")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid Stacks address length")
}

func TestStacksAddress_Equals(t *testing.T) {
	addr1, _ := NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")
	addr2, _ := NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")
	addr3, _ := NewStacksAddress("ST2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7")

	assert.True(t, addr1.Equals(addr2))
	assert.False(t, addr1.Equals(addr3))
}

func TestStacksAddress_IsTestnet(t *testing.T) {
	testnetAddr, _ := NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")
	mainnetAddr, _ := NewStacksAddress("SP2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKQ9H6DPR")

	assert.True(t, testnetAddr.IsTestnet())
	assert.False(t, mainnetAddr.IsTestnet())
}

func TestStacksAddress_IsMainnet(t *testing.T) {
	testnetAddr, _ := NewStacksAddress("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM")
	mainnetAddr, _ := NewStacksAddress("SP2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKQ9H6DPR")

	assert.False(t, testnetAddr.IsMainnet())
	assert.True(t, mainnetAddr.IsMainnet())
}
