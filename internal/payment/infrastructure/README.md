[← payment](../README.md) · **infrastructure** · [root](../../../README.md)

# Infrastructure

> External adapters implementing domain ports (HTTP, blockchain).

## Contents

| Item | Purpose |
|------|---------|
| [`http/`](./http/) | Echo HTTP handlers and request/response DTOs |
| [`blockchain/`](./blockchain/) | Stacks blockchain client adapter |

## Relationships

- **Implements**: Interfaces defined in `../application/command/`
- **Depends on**: `../domain/` for value objects and services

---
*[View on main](https://github.com/x402stacks/stacks-facilitator/tree/main/internal/payment/infrastructure) · Updated: 2025-01-07*
