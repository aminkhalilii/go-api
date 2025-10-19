package mysql

import "go-api/internal/models"

type MysqlRepository struct {
}

func NewMysqlRepository() *MysqlRepository {
	return &MysqlRepository{}
}

func (msql *MysqlRepository) GetAllUsers() ([]models.User, error) {
	users := []models.User{
		{ID: 1, Name: "amin", Email: "aminkhalili@gmail.com", Password: "12345"},
		{ID: 2, Name: "sara", Email: "sara@example.com", Password: "54321"},
	}
	return users, nil

}
