package repositories

import (
	"storeApiRest/database"
	"storeApiRest/models"
)

func CreateOrder(order models.Order) error {
	queryOrder := `INSERT INTO orders(user_id, order_address) VALUES ($1, $2) RETURNING order_id`
	queryOrderProduct := `INSERT INTO order_products(order_id, product_id) VALUES ($1, $2)`
	db := database.GetConnection()
	defer db.Close()

	var orderID int
	err := db.QueryRow(queryOrder, order.UserID.UserID, order.OrderAddress).Scan(&orderID)
	if err != nil {
		return err
	}

	stmtOrderProduct, err := db.Prepare(queryOrderProduct)
	if err != nil {
		return err
	}
	defer stmtOrderProduct.Close()
	for _, product := range order.Products {
		_, err := stmtOrderProduct.Exec(orderID, product)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadOrders() (models.Orders, error) {
	var orders models.Orders
	query := `SELECT o.order_id, o.order_address, o.order_created_at,
    			u.user_id, u.user_full_name,
    			ARRAY_AGG(p.product_id) AS products,
    			 SUM(op.quantity) AS total_quantity,
    			 SUM(op.quantity * p.product_price) AS total_price
				FROM orders o
         			JOIN order_products op ON o.order_id = op.order_id
         			JOIN products p ON op.product_id = p.product_id
         			JOIN users u ON o.user_id = u.user_id
						GROUP BY o.order_id, u.user_id`
	db := database.GetConnection()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order models.Order
		var user models.User
		err := rows.Scan(
			&order.OrderID, &order.OrderAddress, &order.OrderCreatedAt,
			&user.UserID, &user.UserFullName,
			&order.Products,
			&order.TotalQuantity,
			&order.TotalPrice,
		)
		if err != nil {
			return nil, err
		}
		order.UserID = user
		orders = append(orders, &order)
	}
	return orders, nil
}
