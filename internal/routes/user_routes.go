package routes

import (
	"go-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userController := controllers.UserController{}
	r.GET("/", userController.GetAllUsers)

}
