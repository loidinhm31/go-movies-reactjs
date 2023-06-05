package model

import (
	"database/sql"
	"time"
)

type Payment struct {
	ID                uint `gorm:"primaryKey"`
	UserID            uint
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
