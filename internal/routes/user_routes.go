package routes

import (
	"go-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, uc *controllers.UserController) {

	r.GET("/users", uc.GetAllUsers)
	r.GET("/users/:id", uc.GetUserByID)
	r.POST("/users", uc.CreateUser)
	r.PUT("/users/:id", uc.UpdateUser)
	r.DELETE("/users/:id", uc.DeleteUser)

}
