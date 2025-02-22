package main

import (
	"go-lang-role-based-authentication/pkg/database"
	"log"
)

func main() {
	db := database.Connect()
	if db == nil {
		log.Fatal("Failed to connect to the database")
	}
}
