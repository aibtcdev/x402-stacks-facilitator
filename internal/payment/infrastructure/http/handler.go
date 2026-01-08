package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/x402stacks/stacks-facilitator/internal/payment/application/command"
)

// VerifyPaymentHandler interface for verify payment use case
type VerifyPaymentHandler interface {
	Handle(ctx context.Context, cmd command.VerifyPaymentCommand) (command.VerifyPaymentResult, error)
}

// SettlePaymentHandler interface for settle payment use case
type SettlePaymentHandler interface {
	Handle(ctx context.Context, cmd command.SettlePaymentCommand) (command.SettlePaymentResult, error)
}

// Handler handles HTTP requests for payments
type Handler struct {
	verifyHandler VerifyPaymentHandler
	settleHandler SettlePaymentHandler
}

// NewHandler creates a new Handler
func NewHandler(verifyHandler VerifyPaymentHandler, settleHandler SettlePaymentHandler) *Handler {
	return &Handler{
		verifyHandler: verifyHandler,
		settleHandler: settleHandler,
	}
}

// Verify handles POST /api/v1/verify
func (h *Handler) Verify(c echo.Context) error {
	var req VerifyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
	}

	// Validate required fields
	if req.TxID == "" || req.ExpectedRecipient == "" || req.Network == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "missing_required_fields",
			Message: "tx_id, expected_recipient, and network are required",
		})
	}

	// Default token type to STX
	if req.TokenType == "" {
		req.TokenType = "STX"
	}

	cmd := command.VerifyPaymentCommand{
		TxID:              req.TxID,
		TokenType:         req.TokenType,
		ExpectedRecipient: req.ExpectedRecipient,
		MinAmount:         req.MinAmount,
		ExpectedSender:    req.ExpectedSender,
		ExpectedMemo:      req.ExpectedMemo,
		Network:           req.Network,
	}

	result, err := h.verifyHandler.Handle(c.Request().Context(), cmd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "verification_failed",
			Message: err.Error(),
		})
	}

	response := VerifyResponse{
		Valid:            result.Valid,
		TxID:             result.TxID,
		SenderAddress:    result.SenderAddress,
		RecipientAddress: result.RecipientAddress,
		Amount:           result.Amount,
		Fee:              result.Fee,
		Nonce:            result.Nonce,
		Status:           result.Status,
		BlockHeight:      result.BlockHeight,
		TokenType:        result.TokenType,
		Memo:             result.Memo,
		Network:          result.Network,
		Errors:           result.Errors,
	}

	return c.JSON(http.StatusOK, response)
}

// Settle handles POST /api/v1/settle
func (h *Handler) Settle(c echo.Context) error {
	var req SettleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
	}

	// Validate required fields
	if req.SignedTransaction == "" || req.ExpectedRecipient == "" || req.Network == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "missing_required_fields",
			Message: "signed_transaction, expected_recipient, and network are required",
		})
	}

	// Default token type to STX
	if req.TokenType == "" {
		req.TokenType = "STX"
	}

	cmd := command.SettlePaymentCommand{
		SignedTransaction: req.SignedTransaction,
		TokenType:         req.TokenType,
		ExpectedRecipient: req.ExpectedRecipient,
		MinAmount:         req.MinAmount,
		ExpectedSender:    req.ExpectedSender,
		Network:           req.Network,
	}

	result, err := h.settleHandler.Handle(c.Request().Context(), cmd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "settlement_failed",
			Message: err.Error(),
		})
	}

	response := SettleResponse{
		Success:          result.Success,
		TxID:             result.TxID,
		SenderAddress:    result.SenderAddress,
		RecipientAddress: result.RecipientAddress,
		Amount:           result.Amount,
		Fee:              result.Fee,
		Status:           result.Status,
		BlockHeight:      result.BlockHeight,
		TokenType:        result.TokenType,
		Network:          result.Network,
		Errors:           result.Errors,
	}

	if !result.Success {
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, response)
}

// Health handles GET /health
func (h *Handler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthResponse{
		Status: "ok",
	})
}

// RegisterRoutes registers the HTTP routes
func (h *Handler) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")
	api.POST("/verify", h.Verify)
	api.POST("/settle", h.Settle)

	e.GET("/health", h.Health)
}
