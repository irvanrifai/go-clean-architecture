package main

import (
	"log"

	"github.com/irvanrifai/go-clean-architecture/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	db := database.DB
}
