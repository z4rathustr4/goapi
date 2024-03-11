package tools

import (
    log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
    AuthToken string
    Username string
}

type CoinDetails struct {
    Coins int64
    Username string
}

type DBInterface interface {
    GetUserLoginDetails(username string) *LoginDetails 
    GetUserCoins(username string) *CoinDetails
    SetupDB() error
}

func NewDB() (*DBInterface, error) {
    var db DBInterface = &mockDB{}

    var err error = db.SetupDB()
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return &db, nil
}
