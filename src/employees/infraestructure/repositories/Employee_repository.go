package repositories

import (
	"database/sql"
	"demo/src/employees/domain/entities"
	"errors"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (repo *EmployeeRepository) GetAllEmployees() ([]entities.Employee, error) {
	query := "SELECT id, name, age FROM employees"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []entities.Employee
	for rows.Next() {
		var employee entities.Employee
		if err := rows.Scan(&employee.Id, &employee.Name, &employee.Age); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

func (repo *EmployeeRepository) SaveEmployee(employee *entities.Employee) error {
	query := "INSERT INTO employees (name, age) VALUES (?, ?)"
	result, err := repo.db.Exec(query, employee.Name, employee.Age)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	employee.Id = int32(id)
	return nil
}

func (repo *EmployeeRepository) DeleteEmployeeById(id int) error {
	query := "DELETE FROM employees WHERE id = ?"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no employee found with the given id")
	}

	return nil
}

func (repo *EmployeeRepository) UpdateEmployeeById(id int, updatedEmployee *entities.Employee) error {
	query := "UPDATE employees SET name = ?, age = ? WHERE id = ?"
	result, err := repo.db.Exec(query, updatedEmployee.Name, updatedEmployee.Age, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no employee found with the given id")
	}

	return nil
}
