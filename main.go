package main

import (
	"app/config"
	"app/handlers"
	"app/models"
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

	userHandler := &handlers.UserHandler{DB: config.DB}

	router := gin.Default()

	// Rute contoh
	router.GET("/users", userHandler.GetUsers)
	router.POST("/users", userHandler.CreateUser)

	router.Run(":8080")
}
