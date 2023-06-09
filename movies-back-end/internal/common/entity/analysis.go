package entity

type GenreCount struct {
	Name      string
	TypeCode  string
	NumMovies uint
}

type MovieCount struct {
	Year       string
	Month      string
	NumMovies  uint
	Cumulative uint
}

type ViewCount struct {
	Year       string
	Month      string
	NumViewers uint
	Cumulative uint
}
