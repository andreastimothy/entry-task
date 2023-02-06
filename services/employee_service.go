package services

import (
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/dtos"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/models"
	r "git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/repositories"
)

type EmployeeService interface {
	GetAllEmployeees() ([]*models.Employee, error)
	AddEmployee(newEmployee *models.Employee) (*dtos.Employee, error)
	UpdateEmployee(id int, newEmployee *models.Employee) (*dtos.Employee, error)
	DeleteEmployee(id int) (*dtos.EmployeeOutput, error)
}

type employeeService struct {
	employeeRepository r.EmployeeRepository
}

type ESConfig struct {
	EmployeeRepository r.EmployeeRepository
}

func NewEmployeeService(c *ESConfig) EmployeeService {
	return &employeeService{
		employeeRepository: c.EmployeeRepository,
	}
}

func (e *employeeService) GetAllEmployeees() ([]*models.Employee, error) {
	res, err := e.employeeRepository.GetAllEmployeees()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *employeeService) AddEmployee(newEmployee *models.Employee) (*dtos.Employee, error) {
	res, err := e.employeeRepository.PostEmployee(newEmployee)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *employeeService) UpdateEmployee(id int, newEmployee *models.Employee) (*dtos.Employee, error) {
	res, err := e.employeeRepository.UpdateEmployee(id, newEmployee)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *employeeService) DeleteEmployee(id int) (*dtos.EmployeeOutput, error) {
	res, err := e.employeeRepository.DeleteEmployee(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}