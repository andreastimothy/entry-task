package repositories

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/dtos"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/helpers"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EmployeeRepository interface {
	GetAllEmployeees() ([]*models.Employee, error)
	PostEmployee(Employee *models.Employee) (*dtos.Employee, error)
	UpdateEmployee(id int, Employee *models.Employee) (*dtos.Employee, error)
	DeleteEmployee(id int) (*dtos.EmployeeOutput, error)
}

type employeeRepository struct {
	db *gorm.DB
}

type ERConfig struct {
	DB *gorm.DB
}

func NewEmployeeRepository(c *ERConfig) EmployeeRepository {
	return &employeeRepository{
		db: c.DB,
	}
}

func (e *employeeRepository) GetAllEmployeees() ([]*models.Employee, error) {
	var Employee []*models.Employee

	if err := e.db.Raw("SELECT * FROM employees WHERE deleted_at IS NULL ORDER BY id").Scan(&Employee).Error; err != nil {
		return nil, helpers.FailedResponse(http.StatusInternalServerError, "Failed to Fetch All Employeees")
	}
	return Employee, nil
}

func (e *employeeRepository) GetEmployeeById(EmployeeID int) (*models.Employee, error) {
	var Employee *models.Employee
	if err := e.db.Where("id = ?", EmployeeID).First(&Employee).Error; err != nil {
		return nil, helpers.FailedResponse(http.StatusBadRequest, "Failed to Fetch Employee")
	}
	return Employee, nil
}

func (e *employeeRepository) PostEmployee(Employee *models.Employee) (*dtos.Employee, error) {
	if err := e.db.Create(&Employee).Error; err != nil {
		return nil, helpers.FailedResponse(http.StatusInternalServerError, "Failed To Add Employee")
	}
	return &dtos.Employee{
		Name: Employee.Name,
		JobDescription: Employee.JobDescription,
		EntryDate: Employee.EntryDate,
	}, nil
}

func (e *employeeRepository) UpdateEmployee(id int, Employee *models.Employee) (*dtos.Employee, error) {
	_, err := e.GetEmployeeById(id)
	if err != nil {
		return nil, helpers.FailedResponse(http.StatusBadRequest, "Target Employee Not Found")
	}

	if err := e.db.Clauses(clause.Returning{}).Where("id = ?", id).Updates(Employee).Error; err != nil {
		return nil, helpers.FailedResponse(http.StatusInternalServerError, "Failed to Update Employee")
	}
	return &dtos.Employee{
		Name: Employee.Name,
		JobDescription: Employee.JobDescription,
		EntryDate: Employee.EntryDate,
	}, nil
}

func (e *employeeRepository) DeleteEmployee(id int) (*dtos.EmployeeOutput, error) {
	_, err := e.GetEmployeeById(id)
	if err != nil {
		return nil, helpers.FailedResponse(http.StatusBadRequest, "Target Employee Not Found")
	}

	var employee *models.Employee

	if err = e.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&employee).Error; err != nil {
		return nil, helpers.FailedResponse(http.StatusInternalServerError, "Failed to Delete Employee")
	}
	return &dtos.EmployeeOutput{
		Name: employee.Name,
		JobDescription: employee.JobDescription,
		EntryDate: employee.EntryDate,
	}, nil
}