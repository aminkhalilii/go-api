package services

import (
	"go-api/internal/models"
	"go-api/internal/repositories"
)

func GetAllUsers() []models.User {
	return repositories.GetAllUsers()
}
