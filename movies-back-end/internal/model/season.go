package model

import (
	"time"
)

type Season struct {
	ID          int `gorm:"primary_key"`
	Name        string
	AirDate     time.Time
	Description string
	MovieID     int
	Movie       *Movie `gorm:"foreignKey:MovieID"`
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
}

type Episode struct {
	ID        int `gorm:"primary_key"`
	Name      string
	AirDate   time.Time
	Runtime   int
	VideoPath string
	SeasonID  int
	Season    *Season `gorm:"foreignKey:SeasonID"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
