[← domain](../README.md) · **service** · [root](../../../../README.md)

# Service

> Domain services implementing core business operations.

## Contents

| Item | Purpose |
|------|---------|
| [`verification_service.go`](./verification_service.go) | Transaction validation against criteria |
| [`verification_service_test.go`](./verification_service_test.go) | Tests for verification logic |

## Key Types

- `VerificationService` - Validates blockchain transactions
- `BlockchainTransaction` - Domain representation of a tx
- `VerificationCriteria` - Rules for validation (recipient, amount, etc.)
- `VerificationResult` - Valid/invalid with error list

## Relationships

- **Depends on**: `../valueobject/` for domain primitives
- **Consumed by**: `../../application/command/` handlers

---
*[View on main](https://github.com/x402stacks/stacks-facilitator/tree/main/internal/payment/domain/service) · Updated: 2025-01-07*
