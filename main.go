package main

import (
	"app/config"
	"app/models"
	"app/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})

	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
