package routes

import (
	"app/config"
	"app/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	userHandler := &handlers.UserHandler{DB: config.DB}

	// Rute untuk User
	router.GET("/users", userHandler.GetUsers)
	router.POST("/users", userHandler.CreateUser)
}
