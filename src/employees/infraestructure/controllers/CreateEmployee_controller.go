package infraestructure

import (
	"demo/src/employees/application"
	"demo/src/employees/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateEmployeeRequest struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type CreateEmployeeController struct {
	cp application.CreateEmployee
}

func NewCreateEmployeeController(cp application.CreateEmployee) *CreateEmployeeController {
	return &CreateEmployeeController{cp: cp}
}

func (cer *CreateEmployeeController) Execute(c *gin.Context)  { 
	var req CreateEmployeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	employee := entities.Employee{
		Name: req.Name,
		Age:  req.Age,
	}
	
	err := cer.cp.Execute(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee created successfully", "employee": employee})
}
