package main

import (
	"math/rand/v2"
	"time"
)

type LoginRequest struct {
	Number int64 `json:"number"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Number int64 `json:"number"`
	Token string `json:"token"`
}

type TransferRequest struct {
	ToAccount   int   `json:"toAccount"`
	AmountCents int64 `json:"amountCents"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password string `json:"password"`
}

type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	NumberAccount int64     `json:"numberAccount"`
	EncryptedPassword string `json:"-"`
	Balance       int64     `json:"balance"`
	CreatedAt     time.Time `json:"createdAt"`
}

func (a *Account) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}

// TO DO: install bcrypt 
func NewAccount(FirstName, LastName, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte (password), bcrypt.DefaultCost) 
	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:     FirstName,
		LastName:      LastName,
		EncryptedPassword: string(encpw),
		NumberAccount: int64(rand.IntN(1000000)),
		CreatedAt:     time.Now().UTC(),
	}, nil
}

