package controllers

import (
	"go-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users := services.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"users": users})
}
