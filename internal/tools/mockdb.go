package tools

import (
    "time"
)

type mockDB struct{}

var mockLoginData = map[string]LoginDetails {
    "alex" : {
        AuthToken: "ASDF1234",
        Username: "alex",
    },
    "josu" : {
        AuthToken: "FASD4321",
        Username: "josu",
    },
    "mari" : {
        AuthToken: "SDFA1234",
        Username: "mari",
    },
    "juana" : {
        AuthToken: "AFDS4321",
        Username: "juana",
    },
}

var mockCoinDetails = map[string]CoinDetails {
    "alex" : {
        Coins: 100,
        Username: "alex",
    },
    "josu" : {
        Coins: 200,
        Username: "josu",
    },
    "mari" : {
        Coins: 300,
        Username: "mari",
    },
    "juana" : {
        Coins: 400,
        Username: "juana",
    },
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails  {
    time.Sleep(time.Second * 1) 
    
    var clientData = LoginDetails{}
    clientData, ok := mockLoginData[username]
    if !ok {
        return nil
    }

    return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
    time.Sleep(time.Second * 1)
    
    var clientData = CoinDetails{}
    clientData, ok := mockCoinDetails[username]
    if !ok {
        return nil
    }

    return &clientData
}

func (d *mockDB) SetupDB() error {
    return nil
}
