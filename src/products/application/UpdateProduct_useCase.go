package application

import (
	"demo/src/products/domain/entities"
	domain "demo/src/products/domain/interfaces"
)

type UpdateProduct struct {
	productRepository domain.IProduct
}

func NewUpdateProduct(repo domain.IProduct) *UpdateProduct {
	return &UpdateProduct{productRepository: repo}
}

func (up *UpdateProduct) Execute(id int, updatedProduct *entities.Product) error {
	err := up.productRepository.UpdateById(id, updatedProduct)
	if err != nil {
		return err
	}
	return nil
}
