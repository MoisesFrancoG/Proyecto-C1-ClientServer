package application

import (
	"demo/src/products/domain/entities"
	domain "demo/src/products/domain/interfaces"
)

type GetProductById struct {
	productRepository domain.IProduct
}

func NewGetProductById(repo domain.IProduct) *GetProductById {
	return &GetProductById{productRepository: repo}
}

func (gp *GetProductById) Execute(id int) (*entities.Product, error) {
	return gp.productRepository.GetById(id)
}
