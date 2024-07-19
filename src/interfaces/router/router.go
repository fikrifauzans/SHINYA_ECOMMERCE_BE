package router

import (
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

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", userHandler.GetUsers)
		v1.POST("/users", userHandler.CreateUser)
		v1.GET("/users/:id", userHandler.GetUser)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.DELETE("/users/:id", userHandler.DeleteUser)
	}

	return r
}
