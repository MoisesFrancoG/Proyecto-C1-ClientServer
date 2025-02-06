package application

import (
	"demo/src/products/domain/entities"
	domain "demo/src/products/domain/interfaces"
)

type CreateProduct struct {
	productRepository domain.IProduct
}

func NewCreateProduct(repo domain.IProduct) CreateProduct {
	return CreateProduct{productRepository: repo}
}

func (cp *CreateProduct) Execute(product *entities.Product) error {
	return cp.productRepository.SaveProductWithParams(product)
}
