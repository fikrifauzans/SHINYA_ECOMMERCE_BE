package api

import (
	"app/src/domain/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService product.ProductService
}

func NewProductHandler(u product.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: u,
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("perPage", "10"))

	products, total, err := h.productService.GetProducts(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	meta := map[string]interface{}{
		"page":      page,
		"perPage":   perPage,
		"total":     total,
		"totalData": len(products),
		"totalPage": (total + int64(perPage) - 1) / int64(perPage),
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Get List Products Successfully", "data": products, "meta": meta})
}

func (h *ProductHandler) CreateProducts(c *gin.Context) {
	var product product.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	product, err := h.productService.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"error": false, "message": "Product created successfully", "data": product})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Get Product Successfully", "data": product})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product product.Product

	existingProduct, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	product.ID = existingProduct.ID
	product, err = h.productService.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Product updated successfully", "data": product})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.productService.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Product deleted successfully"})
}
