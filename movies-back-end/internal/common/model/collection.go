package model

import (
	"database/sql"
	"time"
)

type PaymentProvider string

const (
	Undefined PaymentProvider = ""
	STRIPE                    = "STRIPE"
)

type Collection struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	MovieID   sql.NullInt64 `gorm:"type:integer, default:null"`
	EpisodeID sql.NullInt64 `gorm:"type:integer, default:null"`
	PaymentID uint
	TypeCode  string
	CreatedAt time.Time
	CreatedBy string
}

type CollectionDetail struct {
	Username    string
	MovieID     uint
	EpisodeID   uint
	TypeCode    string
	Title       string
	SeasonName  string
	EpisodeName string
	ReleaseDate time.Time
	ImageUrl    string
	Description string
	Amount      float64
	CreatedAt   time.Time
}
