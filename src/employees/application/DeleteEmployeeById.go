package application

import domain "demo/src/employees/domain/interfaces"

type DeleteEmployeeById struct {
	employeeRepository domain.IEmployee
}

func NewDeleteEmployeeById(employeeRepository domain.IEmployee) *DeleteEmployeeById {
	return &DeleteEmployeeById{employeeRepository: employeeRepository}
}

func (de *DeleteEmployeeById) Execute(id int) error {
	err := de.employeeRepository.DeleteEmployeeById(id)
	if err != nil {
		return err
	}
	return nil
}
