package models

import (
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Employees   []Employee `json:"employees" gorm:"foreignKey:PositionID"`
}
