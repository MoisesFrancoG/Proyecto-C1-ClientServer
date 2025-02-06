package infraestructure

import (
	"demo/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProductByIdController struct {
	useCase_gpid *application.GetProductById
}

func NewGetProductByIdController(useCase_gpid *application.GetProductById) *GetProductByIdController {
	return &GetProductByIdController{useCase_gpid: useCase_gpid}
}

func (gpc *GetProductByIdController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := gpc.useCase_gpid.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}
