package services

import (
	"errors"
	"karyawan/internal/models"
	"karyawan/internal/repository"
)

func CreateEmployee(employee *models.Employee) error {
	if employee.Name == "" || employee.Age <= 0 || (employee.Gender != "male" && employee.Gender != "female") || employee.PositionID == 0 || employee.DepartmentID == 0 {
		return errors.New("invalid input data")
	}
	return repository.CreateEmployee(employee)
}

func GetEmployeeByID(id uint) (models.Employee, error) {
	return repository.GetEmployeeByID(id)
}

func UpdateEmployee(employee *models.Employee) error {
	if employee.Name == "" || employee.Age <= 0 || (employee.Gender != "male" && employee.Gender != "female") || employee.PositionID == 0 || employee.DepartmentID == 0 {
		return errors.New("invalid input data")
	}
	return repository.UpdateEmployee(employee)
}

func DeleteEmployee(id uint) error {
	return repository.DeleteEmployee(id)
}

func GetAllEmployees() ([]models.Employee, error) {
	return repository.GetAllEmployees()
}
