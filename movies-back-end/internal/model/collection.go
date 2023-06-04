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
	Username  string
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

type Payment struct {
	ID                uint `gorm:"primaryKey"`
	RefID             uint
	TypeCode          string
	Provider          string
	ProviderPaymentID sql.NullString `gorm:"type:varchar(255), default:null"`
	Amount            float64        `gorm:"type:float"`
	Received          float64        `gorm:"type:float"`
	Currency          string
	PaymentMethod     string
	Status            string
	CreatedAt         time.Time
	CreatedBy         string
}
