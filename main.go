package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account number => ", acc.Number)

	return acc
}

func seedAccounts(store Storage) {
	seedAccount(store, "admin", "user", "password")
	seedAccount(store, "admin2", "user2", "password2")
}

func main() {
	// ./bin/gobackend --seed
	seed := flag.Bool("seed", false, "seed the db")
	// here to make it parse the --seed
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("just before seed, and seed is", *seed)
	if *seed {
		fmt.Println("seeding the db")
		seedAccounts(store)
	}

	server := NewAPIServer(":8000", store)
	server.Run()
}
