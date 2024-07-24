package api

import (
	"app/middlewares"
	"app/src/domain/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService user.UserService
}

func NewAuthHandler(u user.UserService) *AuthHandler {
	return &AuthHandler{
		userService: u,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"error": false, "message": "User registered successfully", "data": user})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	user, err := h.userService.GetUserByEmail(input.Email)
	if err != nil || user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Invalid email or password"})
		return
	}

	token, err := middlewares.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Login successful", "token": token})
}
