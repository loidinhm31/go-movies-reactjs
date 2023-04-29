package models

import "time"

type Movie struct {
	ID          int `gorm:"primary_key"`
	Title       string
	ReleaseDate time.Time
	Runtime     int
	MpaaRating  string
	Description string
	Image       string
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
	Genres      []*Genre `gorm:"many2many:movies_genres;"`
}

type Genre struct {
	ID        int `gorm:"primary_key"`
	Genre     string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	Movie     []*Movie `gorm:"many2many:movies_genres;"`
}
