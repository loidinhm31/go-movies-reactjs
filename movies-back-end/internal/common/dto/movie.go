package dto

import "time"

type MovieDto struct {
	ID          uint        `json:"id,omitempty"`
	Title       string      `json:"title,omitempty"`
	TypeCode    string      `json:"type_code"`
	ReleaseDate time.Time   `json:"release_date,omitempty"`
	Runtime     uint        `json:"runtime,omitempty"`
	MpaaRating  string      `json:"mpaa_rating,omitempty"`
	Description string      `json:"description,omitempty"`
	ImageUrl    string      `json:"image_url,omitempty"`
	VideoPath   string      `json:"video_path,omitempty"`
	Price       float64     `json:"price,omitempty"`
	CreatedAt   time.Time   `json:"-"`
	UpdatedAt   time.Time   `json:"-"`
	Genres      []*GenreDto `json:"genres,omitempty"`

	VoteAverage float32 `json:"vote_average,omitempty"`
	VoteCount   uint    `json:"vote_count,omitempty"`
}

type GenreDto struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	TypeCode  string    `json:"type_code"`
	Checked   bool      `json:"checked,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
