package routes

import (
	infraestructure "demo/src/products/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	CreateProductController *infraestructure.CreateProductController
	GetProductsController   *infraestructure.GetProductsController
	UpdateProductController *infraestructure.UpdateProductController
	DeleteProductController *infraestructure.DeleteProductController
	GetProductById *infraestructure.GetProductByIdController
}

func NewProductRoutes(cpc *infraestructure.CreateProductController, gpc *infraestructure.GetProductsController,upc *infraestructure.UpdateProductController,dpc *infraestructure.DeleteProductController,gbd *infraestructure.GetProductByIdController) *ProductRoutes {
	return &ProductRoutes{
		CreateProductController: cpc,
		GetProductsController: gpc,
		UpdateProductController: upc,
		DeleteProductController: dpc,
		GetProductById: gbd,
	}
}

func (pr *ProductRoutes) SetupRoutes(router *gin.Engine) {
	router.POST("/products", func(c *gin.Context) {
		pr.CreateProductController.Execute(c)
	})

	router.GET("/products", func(c *gin.Context) {
		pr.GetProductsController.Execute(c)
	})
	router.PUT("/products/:id", func(c *gin.Context) {
		pr.UpdateProductController.Execute(c)
	})
	router.DELETE("/products/:id",func(c *gin.Context) {
		pr.DeleteProductController.Execute(c)
	})

	router.GET("/products/:id", func(c *gin.Context) { 
		pr.GetProductById.Execute(c)
	})
}
