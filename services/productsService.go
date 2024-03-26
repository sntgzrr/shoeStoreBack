package services

import (
	"storeApiRest/models"
	"storeApiRest/repositories"
)

func CreateProductService(product models.Product) error {
	if err := repositories.CreateProduct(product); err != nil {
		return err
	}
	return nil
}

func ReadProductsService() (models.Products, error) {
	products, err := repositories.ReadProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func UpdateProductService(product models.Product, productID int) error {
	if err := repositories.UpdateProduct(product, productID); err != nil {
		return err
	}
	return nil
}
