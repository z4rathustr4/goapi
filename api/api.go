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

func writeError(w http.ResponseWriter, message string, code int)  {
    resp := Error{
        Code:    code,
        Message: message,
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)

    json.NewEncoder(w).Encode(resp)
}


var (
    RequestErrorHandler = func(w http.ResponseWriter, err error) {
       writeError(w, err.Error(), http.StatusBadRequest) 
    }
    InternalErrorHandler = func(w http.ResponseWriter)  {
        writeError(w, "An unexpected error occurred.", http.StatusInternalServerError)
    }
)
