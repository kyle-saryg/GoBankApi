package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(s Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	// should never fail (unless computer is out of resources)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Seed acc number => ", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "kcs", "dev", "password")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		// Catastrophic if fails
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		// Catastrophic if fails
		log.Fatal(err)
	}

	// Seeding db
	if *seed {
		fmt.Println("seeding the database")
		seedAccounts(store)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
