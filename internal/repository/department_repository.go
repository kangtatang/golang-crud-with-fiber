package repository

import (
	"karyawan/internal/models"
	"karyawan/pkg/db"
)

func CreateDepartment(department *models.Department) error {
	return db.DB.Create(department).Error
}

func GetDepartmentByID(id uint) (models.Department, error) {
	var department models.Department
	err := db.DB.First(&department, id).Error
	return department, err
}

func UpdateDepartment(department *models.Department) error {
	return db.DB.Save(department).Error
}

func DeleteDepartment(id uint) error {
	return db.DB.Delete(&models.Department{}, id).Error
}

func GetAllDepartments() ([]models.Department, error) {
	var departments []models.Department
	err := db.DB.Find(&departments).Error
	return departments, err
}
