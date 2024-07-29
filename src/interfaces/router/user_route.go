package router

import (
	"app/src/application/middlewares"
	"app/src/interfaces/api"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, userHandler *api.UserHandler) {
	userRoutes := r.Group("/users")
	userRoutes.Use(middlewares.JWTAuthMiddleware())
	{
		userRoutes.GET("/", userHandler.GetUsers)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

}
