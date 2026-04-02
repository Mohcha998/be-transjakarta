package models

import "time"

type VehicleLocation struct {
	ID        uint   `gorm:"primaryKey"`
	VehicleID string `gorm:"index"`
	Latitude  float64
	Longitude float64
	Timestamp time.Time
}
