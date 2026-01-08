[← payment](../README.md) · **domain** · [root](../../../README.md)

# Domain

> Core business logic with no external dependencies (pure Go).

## Contents

| Item | Purpose |
|------|---------|
| [`valueobject/`](./valueobject/) | Immutable domain primitives with validation |
| [`service/`](./service/) | Domain services for business operations |

## Relationships

- **Consumed by**: All other layers depend on domain types
- **Principle**: Domain never imports from application or infrastructure

---
*[View on main](https://github.com/x402stacks/stacks-facilitator/tree/main/internal/payment/domain) · Updated: 2025-01-07*
