package model

import "time"

type Rating struct {
	ID        uint `gorm:"primaryKey"`
	Code      string
	Name      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
