package services

import (
	"errors"
	"karyawan/internal/models"
	"karyawan/internal/repository"
)

func CreatePosition(position *models.Position) error {
	if position.Name == "" {
		return errors.New("position name is required")
	}
	return repository.CreatePosition(position)
}

func GetPositionByID(id uint) (models.Position, error) {
	return repository.GetPositionByID(id)
}

func UpdatePosition(position *models.Position) error {
	if position.Name == "" {
		return errors.New("position name is required")
	}
	return repository.UpdatePosition(position)
}

func DeletePosition(id uint) error {
	return repository.DeletePosition(id)
}

func GetAllPositions() ([]models.Position, error) {
	return repository.GetAllPositions()
}
