package application

import (
	"demo/src/employees/domain/entities"
	domain "demo/src/employees/domain/interfaces"
)

type UpdateEmployeeById struct {
	employeeRepository domain.IEmployee
}

func NewUpdateEmployeeById(employeeRepository domain.IEmployee) *UpdateEmployeeById {
	return &UpdateEmployeeById{employeeRepository: employeeRepository}
}

func (ue *UpdateEmployeeById) Execute(id int, updatedEmployee *entities.Employee) error {
	err:= ue.employeeRepository.UpdateEmployeeById(id, updatedEmployee)
	if err != nil {
		return err
	}
	return nil
}
