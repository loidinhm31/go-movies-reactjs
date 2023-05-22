package models

import "time"

type Genre struct {
	ID        int `gorm:"primary_key"`
	Name      string
	TypeCode  string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	Movie     []*Movie `gorm:"many2many:movies_genres;"`
}
