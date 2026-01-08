[← infrastructure](../README.md) · **http** · [root](../../../../README.md)

# HTTP

> Echo framework HTTP handlers exposing payment API endpoints.

## Contents

| Item | Purpose |
|------|---------|
| [`handler.go`](./handler.go) | HTTP handlers for verify, settle, health |
| [`handler_test.go`](./handler_test.go) | Handler integration tests |
| [`dto.go`](./dto.go) | Request/response data transfer objects |

## Endpoints

- `POST /api/v1/verify` - Verify existing transaction
- `POST /api/v1/settle` - Broadcast and confirm transaction
- `GET /health` - Service health check

## Key Types

- `Handler` - Main HTTP handler struct
- `VerifyRequest/Response` - Verification DTOs
- `SettleRequest/Response` - Settlement DTOs
- `RegisterRoutes()` - Mounts all routes on Echo instance

## Relationships

- **Depends on**: `../../application/command/` handlers
- **Framework**: labstack/echo/v4

---
*[View on main](https://github.com/x402stacks/stacks-facilitator/tree/main/internal/payment/infrastructure/http) · Updated: 2025-01-07*
