package dto

type GenreRequest struct {
	Genres []Genres `json:"genres"`
}

type Genres struct {
	Name     string `json:"name"`
	TypeCode string `json:"type_code"`
}
