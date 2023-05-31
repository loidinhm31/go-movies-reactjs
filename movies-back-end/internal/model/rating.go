package model

import "time"

type Rating struct {
	ID        int `gorm:"primary_key"`
	Code      string
	Name      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
