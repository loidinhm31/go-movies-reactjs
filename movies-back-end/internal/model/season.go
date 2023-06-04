package model

import (
	"database/sql"
	"time"
)

type Season struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	AirDate     time.Time
	Description string
	MovieID     uint
	Movie       *Movie `gorm:"foreignKey:MovieID"`
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
}

type Episode struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	AirDate   time.Time
	Runtime   uint
	VideoPath sql.NullString `gorm:"type:varchar(255), default:null"`
	SeasonID  uint
	Season    *Season         `gorm:"foreignKey:SeasonID"`
	Price     sql.NullFloat64 `gorm:"type:float, default:null"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
