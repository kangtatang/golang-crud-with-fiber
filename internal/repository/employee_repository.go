package repository

import (
	"errors"
	"karyawan/internal/models"
	"karyawan/pkg/db"

	"gorm.io/gorm"
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

func GetEmployeeByName(name string) (*models.Employee, error) {
	// Implementasi pencarian employee berdasarkan nama
	// Contoh menggunakan GORM (jika Anda menggunakan GORM sebagai ORM)
	var employee models.Employee
	result := db.DB.Where("name = ?", name).First(&employee)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // No error, just no employee found
		}
		return nil, result.Error // Error during database query
	}
	return &employee, nil
}
