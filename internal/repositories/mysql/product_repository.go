package mysql

import (
	"database/sql"
	"go-api/config"
	"go-api/internal/models"
)

type ProductMysqlRepository struct {
}

func NewProductMysqlRepository() *ProductMysqlRepository {
	return &ProductMysqlRepository{}
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

func (msql *ProductMysqlRepository) GetProductByID(id int) (*models.Product, error) {
	var product models.Product

	row := config.DB.QueryRow("SELECT * FROM products WHERE id=?", id)
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil //if id not exist return nil
		}
		return nil, err
	}

	return &product, nil
}
func (msql *ProductMysqlRepository) CreateProduct(product *models.Product) (*models.Product, error) {

	result, err := config.DB.Exec("insert into products (name,description,price ) values (?,?,?)", product.Name, product.Description, product.Price)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	product.ID = int(id)
	return product, nil
}

func (msql *ProductMysqlRepository) UpdateProduct(id int, product *models.Product) (*models.Product, error) {
	_, err := config.DB.Exec("UPDATE products SET name=?, description=?, price=? WHERE id=?", product.Name, product.Description, product.Price, id)
	if err != nil {
		return nil, err
	}
	return product, nil

}

func (msql *ProductMysqlRepository) DeleteProduct(id int) error {
	_, err := config.DB.Exec("delete from products  where id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil //if id not exist return nil
		}
		return err
	}
	return nil

}
