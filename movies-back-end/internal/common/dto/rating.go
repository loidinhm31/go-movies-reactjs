package dto

import "time"

type RatingDto struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	CreatedBy string    `json:"-"`
	UpdatedAt time.Time `json:"-"`
	UpdatedBy string    `json:"-"`
}
