package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name         string     `json:"name"`
	Age          int        `json:"age"`
	Gender       string     `json:"gender"`
	PositionID   uint       `json:"position_id"`
	DepartmentID uint       `json:"department_id"`
	Position     Position   `json:"position" gorm:"foreignKey:PositionID"`
	Department   Department `json:"department" gorm:"foreignKey:DepartmentID"`
}
