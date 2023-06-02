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
	Username  string   `gorm:"primaryKey"`
	MovieID   uint     `gorm:"primaryKey"`
	Payment   *Payment `gorm:"foreignKey:PaymentID"`
	PaymentID uint
	CreatedAt time.Time
	CreatedBy string
}

type CollectionDetail struct {
	Username    string `gorm:"primaryKey"`
	MovieID     uint   `gorm:"primaryKey"`
	Title       string
	ReleaseDate time.Time
	ImageUrl    string
	Description string
	Amount      float64
	CreatedAt   time.Time
}

type Payment struct {
	ID                uint `gorm:"primaryKey"`
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
