package infraestructure

import (
	"demo/src/products/application"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type LongPollingController struct {
	gp       application.GetProducts
	lastID   int
	mutex    sync.Mutex
	waitList []chan struct{}
}

func NewLongPollingController(gp application.GetProducts) *LongPollingController {
	return &LongPollingController{gp: gp, lastID: 0, waitList: make([]chan struct{}, 0)}
}

func (lp *LongPollingController) CheckNewProducts(c *gin.Context) {
	ch := make(chan struct{})
	lp.mutex.Lock()
	lp.waitList = append(lp.waitList, ch)
	lp.mutex.Unlock()

	select {
	case <-ch:
		c.JSON(http.StatusOK, gin.H{"message": "New product available"})
	case <-time.After(5 * time.Second):
		c.JSON(200, gin.H{"message": "No new products"})
	}
}

func (lp *LongPollingController) NotifyNewProduct() {
	lp.mutex.Lock()
	for _, ch := range lp.waitList {
		close(ch)
	}
	lp.waitList = make([]chan struct{}, 0)
	lp.mutex.Unlock()
}

