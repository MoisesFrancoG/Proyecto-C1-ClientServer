package application

import (
	"demo/src/products/domain/entities"
	domain "demo/src/products/domain/interfaces"
)

type GetProducts struct {
	productRepository domain.IProduct
}

func NewGetProducts(repo domain.IProduct) *GetProducts {
	return &GetProducts{productRepository: repo}
}

func (gp *GetProducts) Execute() ([]entities.Product,error) { //diccionario
	res,err := gp.productRepository.GetAllProducts()
	if err != nil {
		return res,err
	}
	return res,nil
}
