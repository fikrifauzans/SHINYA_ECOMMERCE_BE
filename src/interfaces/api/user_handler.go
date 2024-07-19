package api

import (
	"app/src/domain/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(u user.UserService) *UserHandler {
	return &UserHandler{
		userService: u,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("perPage", "10"))

	users, total, err := h.userService.GetUsers(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	meta := map[string]interface{}{
		"page":      page,
		"perPage":   perPage,
		"total":     total,
		"totalData": len(users),
		"totalPage": (total + int64(perPage) - 1) / int64(perPage),
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Get List Questionnaire Successfully", "data": users, "meta": meta})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
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

	c.JSON(http.StatusCreated, gin.H{"error": false, "message": "User created successfully", "data": user})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Get User Successfully", "data": user})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user user.User

	existingUser, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	user.ID = existingUser.ID
	user, err = h.userService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "User updated successfully", "data": user})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "User deleted successfully"})
}
