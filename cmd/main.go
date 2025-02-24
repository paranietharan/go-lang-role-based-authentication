package main

import (
	"go-lang-role-based-authentication/pkg/database"
	"go-lang-role-based-authentication/pkg/router"
	"log"
)

func main() {
	db := database.Connect()
	if db == nil {
		log.Fatal("Failed to connect to the database")
	}

	router.StartServer()
}
