package repository

import (
	"go-app/internal/models"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) GetLast(id string) (*models.VehicleLocation, error) {
	var data models.VehicleLocation
	err := r.db.Where("vehicle_id = ?", id).Last(&data).Error
	return &data, err
}

func (r *Repo) GetHistory(id string, start, end int64) ([]models.VehicleLocation, error) {
	var list []models.VehicleLocation
	err := r.db.Where("vehicle_id = ? AND extract(epoch from timestamp) BETWEEN ? AND ?", id, start, end).
		Find(&list).Error
	return list, err
}
