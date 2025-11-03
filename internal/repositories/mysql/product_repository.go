package mysql

import (
	"database/sql"
	"go-api/config"
	"go-api/internal/models"
)

type ProductMysqlRepository struct {
}

func NewProductMysqlRepository() *MysqlRepository {
	return &MysqlRepository{}
}

func (msql *ProductMysqlRepository) GetAllProducts() ([]models.Product, error) {
	rows, err := config.DB.Query("select * from products")
	if err != nil {
		return nil, err

	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil

}
func (msql *ProductMysqlRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User

	row := config.DB.QueryRow("SELECT * FROM users WHERE id=?", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil //if id not exist return nil
		}
		return nil, err
	}

	return &user, nil
}
func (msql *ProductMysqlRepository) CreateUser(user *models.User) (*models.User, error) {

	result, err := config.DB.Exec("insert into users (name,email,password ) values (?,?,?)", user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	user.ID = int(id)
	return user, nil
}

func (msql *ProductMysqlRepository) UpdateUser(id int, user *models.User) (*models.User, error) {
	_, err := config.DB.Exec("UPDATE users SET name=?, email=?, password=? WHERE id=?", user.Name, user.Email, user.Password, id)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (msql *ProductMysqlRepository) DeleteUser(id int) error {
	_, err := config.DB.Exec("delete from users  where id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil //if id not exist return nil
		}
		return err
	}
	return nil

}

func (msql *ProductMysqlRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	row := config.DB.QueryRow("SELECT * FROM users WHERE email=?", email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil //if id not exist return nil
		}
		return nil, err
	}

	return &user, nil
}
