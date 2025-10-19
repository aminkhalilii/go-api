package mysql

import (
	"go-api/config"
	"go-api/internal/models"
)

type MysqlRepository struct {
}

func NewMysqlRepository() *MysqlRepository {
	return &MysqlRepository{}
}

func (msql *MysqlRepository) GetAllUsers() ([]models.User, error) {
	rows, err := config.DB.Query("select * from users")
	if err != nil {
		return nil, err

	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}
