package repositories

import (
	"errors"
	"storeApiRest/database"
	"storeApiRest/models"
)

func CreateProduct(product models.Product) error {
	query := `INSERT INTO products(product_name, product_price, product_amount)
				VALUES ($1, $2, $3)`
	db := database.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(product.ProductName, product.ProductPrice, product.ProductAmount)
	if err != nil {
		return err
	}
	i, _ := result.RowsAffected()
	if i != 1 {
		return errors.New("more than 1 row was affected")
	}
	return nil
}

func ReadProducts() (models.Products, error) {
	var products models.Products
	query := `SELECT product_id, product_name, product_price, product_created_at, product_updated_at
				FROM products`
	db := database.GetConnection()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product models.Product
		rows.Scan(
			&product.ProductID,
			&product.ProductName,
			&product.ProductPrice,
			&product.ProductCreatedAt,
			&product.ProductUpdatedAt,
		)
		products = append(products, &product)
	}
	return products, nil
}

func UpdateProduct(product models.Product, productID int) error {
	query := `UPDATE products
				SET product_name = $1, product_price = $2, product_updated_at = now()
					WHERE product_id = $3`
	db := database.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(product.ProductName, product.ProductPrice, productID)
	i, _ := result.RowsAffected()
	if i != 1 {
		return errors.New("more than 1 row was affected")
	}
	return nil
}
