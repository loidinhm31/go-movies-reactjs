package models

import (
	"database/sql"
	"time"
)

type Movie struct {
	ID          int `gorm:"primary_key"`
	Title       string
	ReleaseDate time.Time
	Runtime     int
	MpaaRating  string
	Description string
	ImagePath   sql.NullString `gorm:"type:varchar(255), default:null"`
	VideoPath   sql.NullString `gorm:"type:varchar(255), default:null"`
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

type View struct {
	ID       int `gorm:"primary_key"`
	ViewedBy string
	ViewedAt time.Time
	MovieId  int
}
