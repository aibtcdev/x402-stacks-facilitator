[← root](../README.md) · **internal**

# Internal

> Private application code following Domain-Driven Design architecture.

## Contents

| Item | Purpose |
|------|---------|
| [`payment/`](./payment/) | Payment bounded context (verify/settle use cases) |
| [`stacks/`](./stacks/) | Hiro API client for Stacks blockchain |

## Relationships

- **Consumed by**: `cmd/server/main.go` (entry point)
- **Structure**: DDD layers - domain, application, infrastructure

---
*[View on main](https://github.com/x402stacks/stacks-facilitator/tree/main/internal) · Updated: 2025-01-07*
