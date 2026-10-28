package main

import (
	"math/rand/v2"
	"time"
)

type LoginRequest struct {
	Number int64 `json:"number"`
	Password int64 `json:"password"`
}

type TransferRequest struct {
	ToAccount   int   `json:"toAccount"`
	AmountCents int64 `json:"amountCents"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	NumberAccount int64     `json:"numberAccount"`
	Balance       int64     `json:"balance"`
	CreatedAt     time.Time `json:"createdAt"`
}

func NewAccount(FirstName, LastName string) *Account {
	return &Account{
		FirstName:     FirstName,
		LastName:      LastName,
		NumberAccount: int64(rand.IntN(1000000)),
		CreatedAt:     time.Now().UTC(),
	}
}
