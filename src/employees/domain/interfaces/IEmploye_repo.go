package domain

import "demo/src/employees/domain/entities"

type IEmployee interface {
	GetAllEmployees() ([]entities.Employee, error)
	SaveEmployee(employee *entities.Employee) error
	DeleteEmployeeById(id int) error
	UpdateEmployeeById(id int, updatedEmployee *entities.Employee) error
}