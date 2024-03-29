package models

import (
	"github.com/lib/pq"
	"time"
)

type Order struct {
	OrderID        int           `json:"order_id,omitempty"`
	UserID         User          `json:"user_id,omitempty"`
	Products       pq.Int64Array `json:"products,omitempty"`
	TotalQuantity  int           `json:"total_quantity,omitempty"`
	TotalPrice     int           `json:"total_price,omitempty"`
	OrderAddress   string        `json:"order_address,omitempty"`
	OrderCreatedAt time.Time     `json:"order_created_at,omitempty"`
}

type Orders []*Order
