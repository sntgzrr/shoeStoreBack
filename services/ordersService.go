package services

import (
	"storeApiRest/models"
	"storeApiRest/repositories"
)

func CreateOrderService(order models.Order) error {
	if err := repositories.CreateOrder(order); err != nil {
		return err
	}
	return nil
}

func ReadOrdersService() (models.Orders, error) {
	orders, err := repositories.ReadOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
