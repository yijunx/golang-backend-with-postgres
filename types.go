package main

import "math/rand"

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int
	Balance   int
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(100000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Intn(100000),
		Balance:   0,
	}
}
