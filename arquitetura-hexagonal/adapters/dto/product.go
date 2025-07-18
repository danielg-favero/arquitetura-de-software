package dto

import (
	"github.com/codeedu/go-hexagonal/application"
)

type ProductDTO struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Status string `json:"status"`
}

func NewProductDTO() *ProductDTO {
	return &ProductDTO{}
}

func (productDTO *ProductDTO) Bind(product *application.Product) (*application.Product, error) {
	if productDTO.ID != "" {
		product.ID = productDTO.ID
	}

	product.Name = productDTO.Name
	product.Price = productDTO.Price
	product.Status = productDTO.Status

	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}

	return product, nil
}