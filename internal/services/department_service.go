package services

import (
	"errors"
	"karyawan/internal/models"
	"karyawan/internal/repository"
)

func CreateDepartment(department *models.Department) error {
	if department.Name == "" {
		return errors.New("department name is required")
	}
	return repository.CreateDepartment(department)
}

func GetDepartmentByID(id uint) (models.Department, error) {
	return repository.GetDepartmentByID(id)
}

func UpdateDepartment(department *models.Department) error {
	if department.Name == "" {
		return errors.New("department name is required")
	}
	return repository.UpdateDepartment(department)
}

func DeleteDepartment(id uint) error {
	return repository.DeleteDepartment(id)
}

func GetAllDepartments() ([]models.Department, error) {
	return repository.GetAllDepartments()
}
