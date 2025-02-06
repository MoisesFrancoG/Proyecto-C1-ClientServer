package routes

import (
	infraestructure "demo/src/employees/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type EmployeeRoutes struct {
	CreateEmployeeController *infraestructure.CreateEmployeeController
	GetEmployeesController *infraestructure.GetEmployeesController
	UpdateEmployeeController *infraestructure.UpdateEmployeeController
	DeleteEmployeeController *infraestructure.DeleteEmployeeController
}

func NewEmployeeRoutes(createEmployeeController *infraestructure.CreateEmployeeController,getGetEmployeesController *infraestructure.GetEmployeesController,updateEmployeeController *infraestructure.UpdateEmployeeController, deleteEmployeeController *infraestructure.DeleteEmployeeController) *EmployeeRoutes {
	return &EmployeeRoutes{
		CreateEmployeeController: createEmployeeController,
		GetEmployeesController: getGetEmployeesController,
		UpdateEmployeeController: updateEmployeeController,
		DeleteEmployeeController: deleteEmployeeController,
	}
}

func (em *EmployeeRoutes) SetupRoutes(router *gin.Engine) {
	router.POST("/employees", func(c *gin.Context) {
		em.CreateEmployeeController.Execute(c)
	})

	router.GET("/employees", func(c *gin.Context) {
		em.GetEmployeesController.Execute(c)
	})

	router.PUT("/employees/:id",func(c *gin.Context){
		em.UpdateEmployeeController.Execute(c)
	})

	router.DELETE("employees/:id", func(ctx *gin.Context) {
		em.DeleteEmployeeController.Execute(ctx)
	})
}