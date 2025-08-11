package services

import (
	"encoding/json"
	"os"
	"shopping-cart-app/src/models"
)

type ServicesProdu struct {
}

func NewServiceCar() *ServicesProdu {
	return &ServicesProdu{}
}

func (s *ServicesProdu) GetAllProducts() ([]models.Product, error) {
	data, err := os.ReadFile("src/data/products.json")
	if err != nil {
		return nil, err
	}

	var products []models.Product
	err = json.Unmarshal(data, &products)
	return products, err
}
