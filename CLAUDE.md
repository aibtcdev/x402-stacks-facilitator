# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Stacks Facilitator is a stateless payment verification and settlement service for the Stacks blockchain, implementing the x402 protocol. Built in Go with Domain-Driven Design (DDD) architecture.

## Commands

```bash
# Run tests
go test ./... -v

# Run tests with coverage
go test ./... -cover

# Run specific package tests
go test ./internal/payment/domain/valueobject/... -v

# Download dependencies
go mod download

# Build the binary (note: cmd/server/main.go needs to be created)
go build -o server ./cmd/server

# Run with Docker
docker-compose up -d

# Health check
curl http://localhost:8080/health
```

## Architecture

The codebase follows Domain-Driven Design with clear layer separation:

### Layer Structure

```
internal/
├── payment/                    # Payment bounded context
│   ├── domain/                 # Core business logic (no external dependencies)
│   │   ├── valueobject/        # Immutable value objects (Amount, Network, TokenType, etc.)
│   │   └── service/            # Domain services (VerificationService)
│   ├── application/command/    # Use cases/command handlers
│   └── infrastructure/         # External concerns
│       ├── blockchain/         # StacksClientAdapter
│       └── http/               # Echo HTTP handlers and DTOs
└── stacks/                     # Hiro API client implementation
```

### Key Design Patterns

- **Value Objects** (`domain/valueobject/`): All blockchain primitives (TransactionID, StacksAddress, Amount, Network, TokenType) are immutable value objects with validation in constructors
- **Command Handlers** (`application/command/`): VerifyPaymentHandler and SettlePaymentHandler implement the use cases
- **Port/Adapter**: BlockchainClient and TransactionBroadcaster interfaces in command layer, implemented by StacksClientAdapter
- **Dependency Injection**: Handlers receive dependencies through constructors

### Data Flow

1. HTTP request → `infrastructure/http/Handler` → parses DTO
2. Handler creates Command struct → calls `application/command/*Handler`
3. Command handler uses domain services and infrastructure adapters
4. Results flow back through the same layers

### Token Support

- **STX**: Native token, uses `token_transfer` transaction type
- **sBTC/USDCx**: SIP-010 tokens, uses `contract_call` with `transfer` function

### Network Configuration

Network selection determines Hiro API endpoint:
- Mainnet: `https://api.mainnet.hiro.so`
- Testnet: `https://api.testnet.hiro.so`

## API Endpoints

- `POST /api/v1/verify` - Verify existing blockchain transaction
- `POST /api/v1/settle` - Broadcast signed transaction and wait for confirmation
- `GET /health` - Health check

## Testing Approach

Tests use `github.com/stretchr/testify` for assertions. Each package has `*_test.go` files alongside the implementation. Mock interfaces are defined in command handlers for testing without blockchain access.
