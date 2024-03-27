package models

import "time"

type Order struct {
	OrderID        int       `json:"order_id,omitempty"`
	UserID         User      `json:"user_id,omitempty"`
	Products       Products  `json:"products,omitempty"`
	OrderAddress   string    `json:"order_address,omitempty"`
	OrderCreatedAt time.Time `json:"order_created_at,omitempty"`
}

type Orders []*Order
