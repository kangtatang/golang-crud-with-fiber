package repository

import (
	"karyawan/internal/models"
	"karyawan/pkg/db"
)

func CreateEmployee(employee *models.Employee) error {
	return db.DB.Create(employee).Error
}

func GetEmployeeByID(id uint) (models.Employee, error) {
	var employee models.Employee
	err := db.DB.Preload("Position").Preload("Department").First(&employee, id).Error
	return employee, err
}

func UpdateEmployee(employee *models.Employee) error {
	return db.DB.Save(employee).Error
}

func DeleteEmployee(id uint) error {
	return db.DB.Delete(&models.Employee{}, id).Error
}

func GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	err := db.DB.Preload("Position").Preload("Department").Find(&employees).Error
	return employees, err
}
