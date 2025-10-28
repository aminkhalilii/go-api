package routes

import (
	"go-api/internal/controllers"
	"go-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, uc *controllers.UserController) {

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthRequired())
	{
		authorized.GET("/users", uc.GetAllUsers)
		authorized.GET("/users/:id", uc.GetUserByID)
		authorized.POST("/users", uc.CreateUser)
		authorized.PUT("/users/:id", uc.UpdateUser)
		authorized.DELETE("/users/:id", uc.DeleteUser)

	}

}
