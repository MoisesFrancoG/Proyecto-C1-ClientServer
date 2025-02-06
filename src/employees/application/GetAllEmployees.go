package application

import (
	"demo/src/employees/domain/entities"
	domain "demo/src/employees/domain/interfaces"
)

type GetAllEmployees struct {
	employeeRepository domain.IEmployee
}

func NewGetAllEmployees(employeeRepository domain.IEmployee) *GetAllEmployees {
	return &GetAllEmployees{employeeRepository: employeeRepository}
}

func (gae *GetAllEmployees) Execute() ([]entities.Employee, error) {
	return gae.employeeRepository.GetAllEmployees()
}
