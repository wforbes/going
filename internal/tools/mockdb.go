package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"jimbob": {
		AuthToken: "123ABC",
		Username:  "jimbob",
	},
	"oscar": {
		AuthToken: "456DOG",
		Username:  "oscar",
	},
	"amanda": {
		AuthToken: "789HOT",
		Username:  "amanda",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"jimbob": {
		Coins:    100,
		Username: "jimbob",
	},
	"oscar": {
		Coins:    420,
		Username: "oscar",
	},
	"amanda": {
		Coins:    6969,
		Username: "amanda",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1) // simulating database call

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1) // simulating database call

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
