package models

import "time"

type Product struct {
	ProductID        int       `json:"product_id"`
	ProductName      string    `json:"product_name"`
	ProductPrice     int       `json:"product_price"`
	ProductAmount    int       `json:"product_amount"`
	ProductCreatedAt time.Time `json:"product_created_at"`
	ProductUpdatedAt time.Time `json:"product_updated_at"`
}

type Products []*Product
