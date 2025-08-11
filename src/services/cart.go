package services

import (
	"errors"
	"shopping-cart-app/src/models"
)

type CartService struct {
	cart           models.Cart
	productService *ServicesProdu
}

func NewCartService(productService *ServicesProdu) *CartService {
	return &CartService{
		cart: models.Cart{
			Items: []models.CartItem{},
			Total: 0,
		},
		productService: productService,
	}
}

func (s *CartService) AddProductToCar(idProduct int) error {
	var foundProduct *models.Product

	products, err := s.productService.GetAllProducts()
	if err != nil {
		return err
	}

	for _, product := range products {
		if product.ID == idProduct {
			foundProduct = &product
			break
		}
	}

	if foundProduct == nil {
		return errors.New("product not found")
	}

	for i, item := range s.cart.Items {
		if item.Product.ID == idProduct {
			s.cart.Items[i].Quantity++
			s.calculateTotal()
			return nil
		}
	}

	newItem := models.CartItem{
		Product:  *foundProduct,
		Quantity: 1,
	}

	s.cart.Items = append(s.cart.Items, newItem)
	s.calculateTotal()

	return nil
}

func (s *CartService) GetCart() models.Cart {
	return s.cart
}

func (s *CartService) calculateTotal() {
	total := 0
	for _, item := range s.cart.Items {
		total += item.Product.Price * item.Quantity
	}
	s.cart.Total = total
}
