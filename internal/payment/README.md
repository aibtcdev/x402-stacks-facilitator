[← internal](../README.md) · **payment** · [root](../../README.md)

# Payment

> Bounded context for payment verification and settlement operations.

## Contents

| Item | Purpose |
|------|---------|
| [`domain/`](./domain/) | Core business logic and value objects |
| [`application/`](./application/) | Use cases and command handlers |
| [`infrastructure/`](./infrastructure/) | External adapters (HTTP, blockchain) |

## Relationships

- **Architecture**: Clean Architecture / Hexagonal - domain has no external dependencies
- **Data flow**: HTTP → Application → Domain ← Infrastructure

---
*[View on main](https://github.com/x402stacks/stacks-facilitator/tree/main/internal/payment) · Updated: 2025-01-07*
