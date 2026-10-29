package main

import (
	"log"
	"fmt"
	"log"
)

func seedAccount(fname, lName, pw string, store Storage) *Account {
	acc, err := NewAccount(fname, lName, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", acc.NumberAccount)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount("a", "b", "testing", s)
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the database")
		// seed stuff
		seedAccounts(store)

	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
