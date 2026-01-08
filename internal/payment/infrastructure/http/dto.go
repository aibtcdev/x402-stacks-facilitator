package http

// VerifyRequest represents a verify payment request
type VerifyRequest struct {
	TxID              string  `json:"tx_id"`
	TokenType         string  `json:"token_type,omitempty"`
	ExpectedRecipient string  `json:"expected_recipient"`
	MinAmount         uint64  `json:"min_amount"`
	ExpectedSender    *string `json:"expected_sender,omitempty"`
	ExpectedMemo      *string `json:"expected_memo,omitempty"`
	Network           string  `json:"network"`
}

// VerifyResponse represents a verify payment response
type VerifyResponse struct {
	Valid            bool     `json:"valid"`
	TxID             string   `json:"tx_id"`
	SenderAddress    string   `json:"sender_address"`
	RecipientAddress string   `json:"recipient_address"`
	Amount           uint64   `json:"amount"`
	Fee              uint64   `json:"fee"`
	Nonce            uint64   `json:"nonce,omitempty"`
	Status           string   `json:"status"`
	BlockHeight      uint64   `json:"block_height"`
	TokenType        string   `json:"token_type"`
	Memo             string   `json:"memo,omitempty"`
	Network          string   `json:"network"`
	Errors           []string `json:"errors,omitempty"`
}

// SettleRequest represents a settle payment request
type SettleRequest struct {
	SignedTransaction string  `json:"signed_transaction"`
	TokenType         string  `json:"token_type,omitempty"`
	ExpectedRecipient string  `json:"expected_recipient"`
	MinAmount         uint64  `json:"min_amount"`
	ExpectedSender    *string `json:"expected_sender,omitempty"`
	Network           string  `json:"network"`
}

// SettleResponse represents a settle payment response
type SettleResponse struct {
	Success          bool     `json:"success"`
	TxID             string   `json:"tx_id"`
	SenderAddress    string   `json:"sender_address"`
	RecipientAddress string   `json:"recipient_address"`
	Amount           uint64   `json:"amount"`
	Fee              uint64   `json:"fee"`
	Status           string   `json:"status"`
	BlockHeight      uint64   `json:"block_height"`
	TokenType        string   `json:"token_type"`
	Network          string   `json:"network"`
	Errors           []string `json:"errors,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// HealthResponse represents a health check response
type HealthResponse struct {
	Status string `json:"status"`
}
