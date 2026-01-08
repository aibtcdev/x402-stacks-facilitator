[← application](../README.md) · **command** · [root](../../../../README.md)

# Command

> Command handlers implementing payment verification and settlement use cases.

## Contents

| Item | Purpose |
|------|---------|
| [`verify_payment.go`](./verify_payment.go) | Verify existing blockchain transactions |
| [`verify_payment_test.go`](./verify_payment_test.go) | Tests for verification handler |
| [`settle_payment.go`](./settle_payment.go) | Broadcast and confirm payment transactions |
| [`settle_payment_test.go`](./settle_payment_test.go) | Tests for settlement handler |

## Key Types

- `VerifyPaymentHandler` - Fetches tx, validates against criteria
- `SettlePaymentHandler` - Broadcasts signed tx, waits for confirmation
- `BlockchainClient` - Interface for tx fetching (port)
- `TransactionBroadcaster` - Interface for tx broadcasting (port)

## Relationships

- **Depends on**: `../../domain/service/` for VerificationService
- **Depends on**: `../../domain/valueobject/` for domain primitives
- **Implemented by**: `../../infrastructure/blockchain/` adapters

---
*[View on main](https://github.com/x402stacks/stacks-facilitator/tree/main/internal/payment/application/command) · Updated: 2025-01-07*
