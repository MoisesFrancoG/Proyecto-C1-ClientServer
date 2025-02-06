package domain

import "demo/src/products/domain/entities"

type IProduct interface {
	GetAllProducts()([]entities.Product,error)
	GetById(id int )(*entities.Product, error)
	SaveProductWithParams(product *entities.Product) error
	DeleteById(id int) error
	UpdateById(id int, updatedProduct *entities.Product) error
}
