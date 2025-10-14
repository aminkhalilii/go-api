package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)
	router := gin.Default()
	// Middleware عمومی، مثلا logger یا CORS
	// router.Use(CORSMiddleware())

	return router
}
