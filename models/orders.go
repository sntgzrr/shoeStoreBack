package models

import "time"

type Order struct {
	OrderID        int       `json:"order_id"`
	UserID         User      `json:"user_id"`
	OrderAddress   string    `json:"order_address"`
	OrderCreatedAt time.Time `json:"order_created_at"`
}

type Orders []*Order
