package models

import "time"

type VehicleLocation struct {
	ID        uint      `gorm:"primaryKey"`
	VehicleID string    `gorm:"index;not null" json:"vehicle_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
	CreatedAt time.Time `json:"created_at"`
}
