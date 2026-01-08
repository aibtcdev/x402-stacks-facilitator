[← internal](../README.md) · **stacks** · [root](../../README.md)

# Stacks

> Low-level Hiro API client for Stacks blockchain operations.

## Contents

| Item | Purpose |
|------|---------|
| [`client.go`](./client.go) | HTTP client for Hiro Stacks API |
| [`client_test.go`](./client_test.go) | Client tests with API response parsing |

## Key Types

- `Client` - HTTP client with configurable base URL
- `TransactionResponse` - API response structure for `/extended/v1/tx/{id}`
- `TokenTransferData` - STX native transfer fields
- `ContractCallData` - SIP-010 contract call fields

## API Endpoints Used

- `GET /extended/v1/tx/{txid}` - Fetch transaction details
- `POST /v2/transactions` - Broadcast signed transaction

## Token Parsing

- **STX**: Parsed from `token_transfer` field
- **SIP-010** (sBTC, USDCx): Parsed from `contract_call.function_args`

## Relationships

- **Consumed by**: `../payment/infrastructure/blockchain/` adapter
- **External**: Hiro API (mainnet/testnet)

---
*[View on main](https://github.com/x402stacks/stacks-facilitator/tree/main/internal/stacks) · Updated: 2025-01-07*
