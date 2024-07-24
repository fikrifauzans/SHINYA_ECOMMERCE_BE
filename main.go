package main

import (
	"app/src/infrastructure/config"
	"app/src/interfaces/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()
	// db.Migrate()
	r := router.NewRouter()
	r.Run(":8080")
}
