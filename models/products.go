package models

import "time"

type Product struct {
	ProductID        int       `json:"product_id,omitempty"`
	ProductName      string    `json:"product_name,omitempty"`
	ProductPrice     int       `json:"product_price,omitempty"`
	ProductAmount    int       `json:"product_amount,omitempty"`
	ProductCreatedAt time.Time `json:"product_created_at,omitempty"`
	ProductUpdatedAt time.Time `json:"product_updated_at,omitempty"`
}

type Products []*Product
