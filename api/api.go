package api

// Coin balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin balance response
type CoinBalanceResponse struct {
	// Success code, Usually 200
	Code int

	// Account balance
	Balance int64
}

// Error Response
type Error struct {
	// Error code
	Code int

	// Error Message
	Message string
}
