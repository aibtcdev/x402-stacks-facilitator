[← domain](../README.md) · **valueobject** · [root](../../../../README.md)

# Value Object

> Immutable domain primitives with validation-on-construction.

## Contents

| Item | Purpose |
|------|---------|
| [`amount.go`](./amount.go) | Token amounts in base units (microSTX, satoshis) |
| [`network.go`](./network.go) | Stacks network (mainnet/testnet) with API URLs |
| [`token_type.go`](./token_type.go) | Supported tokens (STX, sBTC, USDCx) |
| [`stacks_address.go`](./stacks_address.go) | Validated Stacks addresses (ST.../SP...) |
| [`transaction_id.go`](./transaction_id.go) | 64-char hex transaction IDs |
| [`payment_status.go`](./payment_status.go) | Payment lifecycle states |

## Design Pattern

All value objects follow the same pattern:
1. Private struct fields
2. `New*()` constructor with validation
3. `String()` for display
4. Domain-specific methods (e.g., `IsMainnet()`, `IsNative()`)

## Relationships

- **Consumed by**: All domain and application code
- **No dependencies**: Pure Go, no external imports

---
*[View on main](https://github.com/x402stacks/stacks-facilitator/tree/main/internal/payment/domain/valueobject) · Updated: 2025-01-07*
