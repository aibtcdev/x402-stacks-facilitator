package valueobject

import (
	"errors"
	"strings"
)

// Network represents a Stacks blockchain network
type Network string

const (
	NetworkMainnet Network = "mainnet"
	NetworkTestnet Network = "testnet"
)

// NewNetwork creates a new Network from a string
func NewNetwork(s string) (Network, error) {
	if s == "" {
		return "", errors.New("network cannot be empty")
	}

	normalized := strings.ToLower(s)
	switch normalized {
	case "mainnet":
		return NetworkMainnet, nil
	case "testnet":
		return NetworkTestnet, nil
	default:
		return "", errors.New("unsupported network: " + s)
	}
}

// String returns the network as a string
func (n Network) String() string {
	return string(n)
}

// APIBaseURL returns the Hiro API base URL for this network
func (n Network) APIBaseURL() string {
	switch n {
	case NetworkMainnet:
		return "https://api.mainnet.hiro.so"
	case NetworkTestnet:
		return "https://api.testnet.hiro.so"
	default:
		return ""
	}
}

// IsMainnet returns true if this is mainnet
func (n Network) IsMainnet() bool {
	return n == NetworkMainnet
}

// IsTestnet returns true if this is testnet
func (n Network) IsTestnet() bool {
	return n == NetworkTestnet
}
