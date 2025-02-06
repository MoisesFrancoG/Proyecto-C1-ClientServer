package dependencies

import (
	"database/sql"
	"demo/src/products/application"
	infraestructure "demo/src/products/infraestructure/controllers"
	"demo/src/products/infraestructure/repositories"
	"demo/src/products/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

type ProductDependencies struct {
	DB *sql.DB
}

func NewProductDependencies(db *sql.DB) *ProductDependencies {
	return &ProductDependencies{DB: db}
}

func (pd *ProductDependencies) Execute(r *gin.Engine) {
	productRepo := repositories.NewProductRepository(pd.DB)
	lpController := infraestructure.NewLongPollingController(*application.NewGetProducts(productRepo))
	createProductCase := application.NewCreateProduct(productRepo)
	createProductController := infraestructure.NewCreateProductController(createProductCase, lpController)

	deleteProductCase := application.NewDeleteProduct(productRepo)
	deleteProductController := infraestructure.NewDeleteProductController(deleteProductCase)

	getAllProductsCase := application.NewGetProducts(productRepo)

	getAllProductsController := infraestructure.NewGetRecentProductsController(*getAllProductsCase)
	getUpdatedProducts := infraestructure.NewGetUpdatedPricesController(*getAllProductsCase)

	updatedProductCase := application.NewUpdateProduct(productRepo)
	updatedProductController := infraestructure.NewUpdateProductController(updatedProductCase)

	getByIdCase := application.NewGetProductById(productRepo)
	getByIdController := infraestructure.NewGetProductByIdController(getByIdCase)


	productRoutes := routes.NewProductRoutes(createProductController, getAllProductsController, updatedProductController, deleteProductController, getByIdController, getUpdatedProducts,lpController)
	productRoutes.SetupRoutes(r)
}