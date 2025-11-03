package routes

import (
	"go-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, uc *controllers.UserController, ac *controllers.AuthController) {

	r.GET("/Products", ac.Register)
	// r.POST("/login", ac.Login)
	// r.POST("/refresh", ac.RefreshToken)
	// authorized := r.Group("/")
	// authorized.Use(middlewares.AuthRequired())
	// {
	// 	authorized.GET("/users", uc.GetAllUsers)
	// 	authorized.GET("/users/:id", uc.GetUserByID)
	// 	authorized.POST("/users", uc.CreateUser)
	// 	authorized.PUT("/users/:id", uc.UpdateUser)
	// 	authorized.DELETE("/users/:id", uc.DeleteUser)

	// }

}
