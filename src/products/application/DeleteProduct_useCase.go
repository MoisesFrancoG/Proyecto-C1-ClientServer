package application

import domain "demo/src/products/domain/interfaces"

type DeleteProduct struct {
	productRepository domain.IProduct
}

func NewDeleteProduct(repo domain.IProduct) *DeleteProduct {
	return &DeleteProduct{productRepository: repo}
}

func (dp *DeleteProduct) Execute(id int) error {
	err := dp.productRepository.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
