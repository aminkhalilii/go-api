package routes

import (
	"go-api/internal/controllers"
	"go-api/internal/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, uc *controllers.UserController, ac *controllers.AuthController) {

	rateLimitGroup := r.Group("/")
	rateLimitGroup.Use(middlewares.RateLimitMiddleware(3, 1*time.Minute))
	{
		rateLimitGroup.POST("/register", ac.Register)
		rateLimitGroup.POST("/login", ac.Login)
		rateLimitGroup.POST("/refresh", ac.RefreshToken)
	}

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
