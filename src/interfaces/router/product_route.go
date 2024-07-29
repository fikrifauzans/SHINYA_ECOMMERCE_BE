package router

import (
	"app/src/interfaces/api"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup, productHandler *api.ProductHandler) {
	productRoutes := r.Group("/product")
	productRoutes.GET("/", productHandler.GetProducts)
	productRoutes.POST("/", productHandler.CreateProducts)
	productRoutes.GET("/:id", productHandler.GetProductByID)
	productRoutes.PUT("/:id", productHandler.UpdateProduct)
	productRoutes.DELETE("/:id", productHandler.DeleteProduct)
}
