[← infrastructure](../README.md) · **blockchain** · [root](../../../../README.md)

# Blockchain

> Adapter connecting domain layer to Stacks blockchain via Hiro API.

## Contents

| Item | Purpose |
|------|---------|
| [`stacks_client_adapter.go`](./stacks_client_adapter.go) | Implements BlockchainClient and TransactionBroadcaster |

## Key Types

- `StacksClientAdapter` - Wraps Stacks client for domain use
  - `GetTransactionWithRetry()` - Fetch tx with retry logic
  - `WaitForConfirmation()` - Poll until confirmed/failed
  - `BroadcastTransaction()` - Submit signed tx to network

## Relationships

- **Implements**: `BlockchainClient`, `TransactionBroadcaster` from application layer
- **Depends on**: `../../../stacks/` for low-level API calls
- **Network routing**: Maintains separate mainnet/testnet clients

---
*[View on main](https://github.com/x402stacks/stacks-facilitator/tree/main/internal/payment/infrastructure/blockchain) · Updated: 2025-01-07*
