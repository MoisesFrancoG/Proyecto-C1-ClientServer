package infraestructure

import (
	"demo/src/products/application"
	"demo/src/products/domain/entities"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GetRecentProductsController struct {
	gp application.GetProducts
	lastCheck time.Time
}

func NewGetRecentProductsController(gp application.GetProducts) *GetRecentProductsController {
	return &GetRecentProductsController{gp: gp, lastCheck: time.Now()}
}

func (rpc *GetRecentProductsController) Execute(c *gin.Context) {
	products, err := rpc.gp.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}

	var recentProducts []entities.Product
	for _, p := range products {
		if time.Now().Sub(rpc.lastCheck) < 30*time.Second {
			recentProducts = append(recentProducts, p)
		}
	}

	rpc.lastCheck = time.Now()
	c.JSON(http.StatusOK, gin.H{"Products": recentProducts})
}
