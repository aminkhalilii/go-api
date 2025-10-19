package routes

import (
	"go-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes(r *gin.Engine, userController *controllers.UserController) {
	RegisterUserRoutes(r, userController)
}
