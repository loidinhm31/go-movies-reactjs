package models

type GenreCount struct {
	Name      string
	TypeCode  string
	NumMovies int
}

type MovieCount struct {
	Year       string
	Month      string
	NumMovies  int
	Cumulative int
}

type ViewCount struct {
	Year       string
	Month      string
	NumViewers int
	Cumulative int
}
