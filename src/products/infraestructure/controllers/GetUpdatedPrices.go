package infraestructure

import (
	"demo/src/products/application"
	"demo/src/products/domain/entities"
	"net/http"
	"github.com/gin-gonic/gin"
)

type GetUpdatedPricesController struct {
	gp             application.GetProducts
	previousPrices map[int32]float32
}

func NewGetUpdatedPricesController(gp application.GetProducts) *GetUpdatedPricesController {
	return &GetUpdatedPricesController{
		gp:             gp,
		previousPrices: make(map[int32]float32),
	}
}

func (upc *GetUpdatedPricesController) Execute(c *gin.Context) {
	products, err := upc.gp.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}

	var updatedProducts []entities.Product
	for _, p := range products {
		if prevPrice, exists := upc.previousPrices[p.Id]; exists {
			if prevPrice != p.Price {
				updatedProducts = append(updatedProducts, p)
			}
		}
		upc.previousPrices[p.Id] = p.Price
	}

	c.JSON(http.StatusOK, gin.H{"updated_products": updatedProducts})
}
