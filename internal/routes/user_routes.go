package routes

import (
	"go-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, uc *controllers.UserController) {

	r.GET("/", uc.GetAllUsers)

}
