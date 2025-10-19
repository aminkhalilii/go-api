package controllers

import (
	"go-api/internal/services"
	"net/http"

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
