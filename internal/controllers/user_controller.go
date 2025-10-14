package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// متد کنترلر برای گرفتن همه‌ی کاربران
func (uc UserController) GetAllUsers(c *gin.Context) {
	users := []map[string]string{
		{"id": "1", "name": "Ali"},
		{"id": "2", "name": "Sara"},
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
