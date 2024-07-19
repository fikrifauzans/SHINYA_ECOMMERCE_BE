package main

import (
	"app/config"
	"app/infrastructure/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	r := router.NewRouter()
	r.Run(":8080")
}
