package controllers

import (
	"go-api/config"
	"go-api/internal/models"
	"go-api/internal/services"
	"go-api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthController struct {
	authService services.AuthServiceInterface
	userService services.UserServiceInterface
}

func NewAuthController(authService services.AuthServiceInterface) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Login(c *gin.Context) {
	var user struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
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
func (ac *AuthController) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "Refresh token required")
		return
	}

	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		utils.Error(c, http.StatusUnauthorized, "Invalid refresh token")
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))

	user, err := ac.userService.GetUserByID(userID)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, "User not found")
		return
	}

	tokens, err := ac.authService.GenerateToken(user)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Could not create new tokens")
		return
	}

	utils.Success(c, http.StatusOK, "Tokens refreshed", gin.H{
		"access_token":  tokens["access_token"],
		"refresh_token": tokens["refresh_token"],
	})
}
