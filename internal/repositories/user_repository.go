package repositories

import (
	"go-api/internal/models"
)

// متد برای گرفتن همه کاربران
func GetAllUsers() []models.User {
	users := []models.User{
		{ID: 1, Name: "amin", Email: "aminkhalili@gmail.com", Password: "12345"},
		{ID: 2, Name: "sara", Email: "sara@example.com", Password: "54321"},
	}
	return users
}
