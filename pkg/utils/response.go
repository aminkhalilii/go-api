package utils

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(200, ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ApiResponse{
		Status:  "error",
		Message: message,
	})
}
