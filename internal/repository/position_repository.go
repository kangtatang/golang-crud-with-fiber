package repository

import (
	"karyawan/internal/models"
	"karyawan/pkg/db"
)

func CreatePosition(position *models.Position) error {
	return db.DB.Create(position).Error
}

func GetPositionByID(id uint) (models.Position, error) {
	var position models.Position
	err := db.DB.First(&position, id).Error
	return position, err
}

func UpdatePosition(position *models.Position) error {
	return db.DB.Save(position).Error
}

func DeletePosition(id uint) error {
	return db.DB.Delete(&models.Position{}, id).Error
}

func GetAllPositions() ([]models.Position, error) {
	var positions []models.Position
	err := db.DB.Find(&positions).Error
	return positions, err
}
