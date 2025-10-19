package mysql

import (
	"database/sql"
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
func (msql *MysqlRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User

	row := config.DB.QueryRow("SELECT * FROM users WHERE id=?", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // کاربری با این ID پیدا نشد
		}
		return nil, err
	}

	return &user, nil
}
func (msql *MysqlRepository) CreateUser(user *models.User) (*models.User, error) {

	result, err := config.DB.Exec("insert into users (name,email,password ) values (?,?,?)", user.Name, user.Email, user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // کاربری با این ID پیدا نشد
		}
		return nil, err
	}
	id, _ := result.LastInsertId()
	user.ID = int(id)
	return user, nil
}
