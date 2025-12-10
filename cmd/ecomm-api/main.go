package main

import (
	"log"

	db "github.com/M-oses340/Microservices-Database-Setup/db/migrations"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("error opening database: %v", err)
	}
	defer db.Close()
	log.Println("Database opened")
}
