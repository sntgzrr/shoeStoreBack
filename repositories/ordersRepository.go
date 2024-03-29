package repositories

import (
	"errors"
	"storeApiRest/database"
	"storeApiRest/models"
	"strconv"
	"strings"
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

	products := strings.Split(order.Products, ",")
	for _, product := range products {
		id, _ := strconv.Atoi(product)
		resultOrderProduct, err := stmtOrderProduct.Exec(order.OrderID, id)
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
    			u.user_id, u.user_full_name,
    			STRING_AGG(CAST(p.product_id AS VARCHAR), ', ') AS products
				FROM orders o
         			JOIN order_products op ON o.order_id = op.order_id
         			JOIN products p ON op.product_id = p.product_id
         			JOIN users u ON o.user_id = u.user_id
						GROUP BY o.order_id, o.order_address, o.order_created_at, u.user_id, u.user_full_name`
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
		)
		if err != nil {
			return nil, err
		}
		order.UserID = user
		orders = append(orders, &order)
	}
	return orders, nil
}
