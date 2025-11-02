package controllers

import (
	"go-api/internal/models"
	"go-api/internal/services"
	"go-api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthServiceInterface
}

func NewAuthController(authService services.AuthServiceInterface) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input")
		return
	}

	userLogin, err := ac.authService.Login(user.Email, user.Password)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := ac.authService.GenerateToken(userLogin)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Could not create token")
		return
	}
	user.Password = ""
	utils.Success(c, http.StatusOK, "Login successful", gin.H{
		"token": token,
		"user":  user,
	})
}
func (ac *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input")
		return
	}

	newUser, err := ac.authService.Register(&user)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	newUser.Password = ""
	utils.Success(c, http.StatusCreated, "User registered successfully", newUser)
}
