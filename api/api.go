package api

import (
    "encoding/json"
    "net/http"
)
// Coin balance params
type CoinBalanceParams struct {
    Username string
}
// Coin balance response
type CoinBalanceResponse struct {
    // success code eg 200
    Code int
    // account balance
    Balance int64
}

// Error Response
type Error struct {
    // error code
    Code int
    // Error msg
    Message string
}
