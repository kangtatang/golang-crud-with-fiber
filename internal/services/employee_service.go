package services

import (
	"errors"
	"karyawan/internal/models"
	"karyawan/internal/repository"
	"regexp"
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
	// Cek apakah nama kosong, kurang dari 3 karakter, atau mengandung karakter selain huruf dan spasi
	if len(employee.Name) < 3 || !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(employee.Name) {
		return errors.New("invalid name input")
	}

	// Cek apakah nama unik kecuali untuk ID yang sama
	existingEmployee, err := repository.GetEmployeeByName(employee.Name)
	if err != nil {
		return err
	}
	if existingEmployee != nil && existingEmployee.ID != employee.ID {
		return errors.New("name must be unique")
	}

	// Cek apakah umur kurang dari atau sama dengan 0
	if employee.Age <= 0 {
		return errors.New("invalid age input")
	}

	// Cek apakah gender valid
	if employee.Gender != "male" && employee.Gender != "female" {
		return errors.New("invalid gender input")
	}

	// Cek apakah PositionID dan DepartmentID valid
	if employee.PositionID == 0 || employee.DepartmentID == 0 {
		return errors.New("PositionID and DepartmentID must be non-zero numbers")
	}

	return repository.UpdateEmployee(employee)
}

func DeleteEmployee(id uint) error {
	return repository.DeleteEmployee(id)
}

func GetAllEmployees() ([]models.Employee, error) {
	return repository.GetAllEmployees()
}
