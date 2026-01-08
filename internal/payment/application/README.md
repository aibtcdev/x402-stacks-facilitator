[← payment](../README.md) · **application** · [root](../../../README.md)

# Application

> Use case orchestration layer implementing CQRS command pattern.

## Contents

| Item | Purpose |
|------|---------|
| [`command/`](./command/) | Command handlers for verify and settle operations |

## Relationships

- **Depends on**: `../domain/` for business logic and value objects
- **Consumed by**: `../infrastructure/http/` handlers invoke commands

---
*[View on main](https://github.com/x402stacks/stacks-facilitator/tree/main/internal/payment/application) · Updated: 2025-01-07*
