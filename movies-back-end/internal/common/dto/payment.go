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
