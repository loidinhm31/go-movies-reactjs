package dto

import "time"

type MovieDto struct {
	ID          int         `json:"id,omitempty"`
	Title       string      `json:"title"`
	ReleaseDate time.Time   `json:"release_date"`
	Runtime     int         `json:"runtime"`
	MpaaRating  string      `json:"mpaa_rating"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	CreatedAt   time.Time   `json:"-"`
	UpdatedAt   time.Time   `json:"-"`
	Genres      []*GenreDto `json:"genres,omitempty"`
}

type GenreDto struct {
	ID        int       `json:"id"`
	Genre     string    `json:"genre"`
	Checked   bool      `json:"checked"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
