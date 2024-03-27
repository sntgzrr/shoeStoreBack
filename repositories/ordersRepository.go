package repositories

import (
	"errors"
	"storeApiRest/database"
	"storeApiRest/models"
)

func CreateOrder(order models.Order) error {
	queryOrder := `INSERT INTO orders(user_id, order_address)
					VALUES ($1, $2)`
	queryOrderProduct := `INSERT INTO order_products(order_id, product_id)
							VALUES ($1, $2)`
	db := database.GetConnection()
	defer db.Close()

	stmtOrder, err := db.Prepare(queryOrder)
	if err != nil {
		return err
	}
	defer stmtOrder.Close()
	resultOrder, err := stmtOrder.Exec(order.UserID.UserID, order.OrderAddress)
	if err != nil {
		return err
	}
	i, _ := resultOrder.RowsAffected()
	if i != 1 {
		return errors.New("more than 1 row was affected")
	}

	stmtOrderProduct, err := db.Prepare(queryOrderProduct)
	if err != nil {
		return err
	}
	defer stmtOrderProduct.Close()
	for _, product := range order.Products {
		resultOrderProduct, err := stmtOrderProduct.Exec(order.OrderID, product.ProductID)
		if err != nil {
			return err
		}
		i, _ = resultOrderProduct.RowsAffected()
		if i != 1 {
			return errors.New("more than 1 row was affected")
		}
	}
	return nil
}

func ReadOrders() (models.Orders, error) {
	var orders models.Orders
	query := `SELECT o.order_id, o.order_address, o.order_created_at, 
				u.user_id, u.user_full_name, u.user_email, u.user_created_at, u.user_updated_at,
				p.product_id, p.product_name, p.product_price, p.product_amount, p.product_created_at, p.product_updated_at
				FROM orders o
					JOIN users u ON o.user_id = u.user_id
					JOIN order_products op ON o.order_id = op.order_id
					JOIN products p ON op.product_id = p.product_id`
	db := database.GetConnection()
	defer db.Close()
	var order models.Order
	var user models.User
	var product models.Product
	if err := db.QueryRow(query).Scan(&order.OrderID, &order.OrderAddress, &order.OrderCreatedAt,
		&user.UserID, &user.UserFullName, &user.UserEmail, &user.UserCreatedAt, &user.UserUpdatedAt,
		&product.ProductID, &product.ProductName, &product.ProductPrice, &product.ProductAmount, &product.ProductCreatedAt, &product.ProductUpdatedAt,
	); err != nil {
		return nil, err
	}
	order.UserID = user
	products, err := ReadProductsByOrderID(order.OrderID)
	if err != nil {
		return nil, err
	}
	order.Products = products
	orders = append(orders, &order)
	return orders, nil
}

func ReadProductsByOrderID(orderID int) (models.Products, error) {
	var products models.Products
	query := `SELECT p.product_id, p.product_name, p.product_price, p.product_amount, p.product_created_at, p.product_updated_at
				FROM products p
					JOIN order_products op ON p.product_id = op.product_id
						WHERE op.order_id = $1`
	db := database.GetConnection()
	defer db.Close()
	var product models.Product
	if err := db.QueryRow(query, orderID).Scan(&product.ProductID, &product.ProductName, &product.ProductPrice, &product.ProductAmount, &product.ProductCreatedAt, &product.ProductUpdatedAt); err != nil {
		return nil, err
	}
	products = append(products, &product)
	return products, nil
}
