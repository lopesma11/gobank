package main

import "math/rand/v2"

type Account struct {
	ID            int    `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	NumberAccount int64  `json:"numberAccount"`
	Balance       int64  `json:"balance"`
}

func NewAccount(FirstName, LastName string) *Account {
	return &Account{
		ID:            rand.IntN(10000),
		FirstName:     FirstName,
		LastName:      LastName,
		NumberAccount: int64(rand.IntN(1000000)),
	}
}
