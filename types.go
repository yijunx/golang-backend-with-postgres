package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Number int64  `json:"number"`
	Token  string `json:"token"`
}

type LoginRequest struct {
	Number   int64  `json:"number"`
	Password string `json:"password"`
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Number            int64     `json:"number"`
	EncryptedPassword string    `json:"-"` // not going to return this...
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (a *Account) ValidatePassword(password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(password)) != nil {
		return false
	} else {
		return true
	}
}

func NewAccount(firstName, lastName, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		// ID:        rand.Intn(100000),
		// we will let postgres to generate the int auto increment id
		// or we can use the uuid
		FirstName:         firstName,
		LastName:          lastName,
		Number:            int64(rand.Intn(10000000)),
		EncryptedPassword: string(encpw),
		// Balance:   0,
		// cos it will be init as 0
		CreatedAt: time.Now().UTC(),
	}, nil
}
