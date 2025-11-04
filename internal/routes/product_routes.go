package routes

import (
	"go-api/internal/controllers"
	"go-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, pc *controllers.ProductController) {

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthRequired())
	{
		authorized.GET("/products", pc.GetAllProducts)
		authorized.GET("/products/:id", pc.GetProductByID)
		authorized.POST("/products", pc.CreateProduct)
		authorized.PUT("/products/:id", pc.UpdateProduct)
		authorized.DELETE("/products/:id", pc.DeleteProduct)

	}

}
