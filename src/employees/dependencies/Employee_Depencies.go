package dependencies

import (
	"database/sql"
	"demo/src/employees/application"
	infraestructure "demo/src/employees/infraestructure/controllers"
	"demo/src/employees/infraestructure/repositories"
	"demo/src/employees/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

type EmployeeDependencies struct {
	DB *sql.DB
}

func NewEmployeeDependencies(db *sql.DB) *EmployeeDependencies {
	return &EmployeeDependencies{DB: db}
}

func (ed *EmployeeDependencies) Execute(r *gin.Engine) {
	employeRepo := repositories.NewEmployeeRepository(ed.DB)

	createEmployeeCase := application.NewCreateEmployee(employeRepo)
	createEmployeeController := infraestructure.NewCreateEmployeeController(createEmployeeCase)

	deleteEmployeeCase := application.NewDeleteEmployeeById(employeRepo)
	deleteEmployeeController := infraestructure.NewDeleteEmployeeController(deleteEmployeeCase)

	getAllEmployeesCase := application.NewGetAllEmployees(employeRepo)
	getAllEmployeesController := infraestructure.NewGetEmployeesController(getAllEmployeesCase)

	updateEmployeCase := application.NewUpdateEmployeeById(employeRepo)
	updateEmployeController := infraestructure.NewUpdateEmployeeController(updateEmployeCase)

	employeeRoutes := routes.NewEmployeeRoutes(createEmployeeController, getAllEmployeesController, updateEmployeController, deleteEmployeeController)
	employeeRoutes.SetupRoutes(r)
}