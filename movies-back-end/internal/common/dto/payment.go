package dto

import (
	"time"
)

type PaymentDto struct {
	ID                uint      `json:"id"`
	UserID            uint      `json:"user_id"`
	RefID             uint      `json:"ref_id"`
	TypeCode          string    `json:"type_code"`
	Provider          string    `json:"provider,omitempty"`
	ProviderPaymentID string    `json:"provider_payment_id,omitempty"`
	Amount            float64   `json:"amount"`
	Currency          string    `json:"currency"`
	PaymentMethod     string    `json:"payment_method"`
	Status            string    `json:"status,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
}

type CustomPaymentDto struct {
	ID            uint      `json:"id"`
	TypeCode      string    `json:"type_code"`
	MovieTitle    string    `json:"movie_title"`
	SeasonName    string    `json:"season_name"`
	EpisodeName   string    `json:"episode_name"`
	Provider      string    `json:"provider"`
	PaymentMethod string    `json:"payment_method"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}
