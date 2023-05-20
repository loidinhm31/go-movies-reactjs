package dto

import "time"

type MovieDto struct {
	ID          int         `json:"id,omitempty"`
	Title       string      `json:"title,omitempty"`
	TypeCode    string      `json:"type_code"`
	ReleaseDate time.Time   `json:"release_date,omitempty"`
	Runtime     int         `json:"runtime,omitempty"`
	MpaaRating  string      `json:"mpaa_rating,omitempty"`
	Description string      `json:"description,omitempty"`
	ImagePath   string      `json:"image_path,omitempty"`
	VideoPath   string      `json:"video_path,omitempty"`
	CreatedAt   time.Time   `json:"-"`
	UpdatedAt   time.Time   `json:"-"`
	Genres      []*GenreDto `json:"genres,omitempty"`

	VoteAverage float32 `json:"vote_average,omitempty"`
	VoteCount   int     `json:"vote_count,omitempty"`
}

type GenreDto struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	TypeCode  string    `json:"type_code"`
	Checked   bool      `json:"checked,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
