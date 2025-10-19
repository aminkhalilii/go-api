package controllers

import (
	"go-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users := uc.userService.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"users": users})
}
func (uc *UserController) GetUserByID(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	user := uc.userService.GetUserByID(id)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
