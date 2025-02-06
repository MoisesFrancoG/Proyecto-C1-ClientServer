package infraestructure

import (
	"demo/src/employees/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteEmployeeController struct {
	useCase_dp *application.DeleteEmployeeById
}

func NewDeleteEmployeeController(useCase_dp *application.DeleteEmployeeById) *DeleteEmployeeController {
	return &DeleteEmployeeController{useCase_dp: useCase_dp}
}

func (dec *DeleteEmployeeController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = dec.useCase_dp.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}