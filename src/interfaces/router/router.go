package router

import (
	usecase "app/src/application/usecase"
	repository "app/src/infrastructure/persistence"
	"app/src/interfaces/api"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	userRepository := repository.NewUserRepository()
	userService := usecase.NewUserUsecase(userRepository)

	productRepository := repository.NewProductRepository()
	productService := usecase.NewProductUsecase(productRepository)

	userHandler := api.NewUserHandler(userService)
	authHandler := api.NewAuthHandler(userService)
	productHandler := api.NewProductHandler(productService)

	v1 := r.Group("/api/v1")
	{
		AuthRoutes(v1, authHandler)
		UserRoutes(v1, userHandler)
		ProductRoutes(v1, productHandler)
	}

	return r
}
