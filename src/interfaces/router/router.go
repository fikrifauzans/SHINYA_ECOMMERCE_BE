package router

import (
	"app/middlewares"
	userRepo "app/repository/user"
	"app/src/interfaces/api"
	userUsecase "app/usecase/user"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	userRepository := userRepo.NewUserRepository()
	userService := userUsecase.NewUserUsecase(userRepository)
	userHandler := api.NewUserHandler(userService)
	authHandler := api.NewAuthHandler(userService)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)

		userRoutes := v1.Group("/users")
		userRoutes.Use(middlewares.JWTAuthMiddleware())
		{
			userRoutes.GET("/", userHandler.GetUsers)
			userRoutes.POST("/", userHandler.CreateUser)
			userRoutes.GET("/:id", userHandler.GetUser)
			userRoutes.PUT("/:id", userHandler.UpdateUser)
			userRoutes.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	return r
}
