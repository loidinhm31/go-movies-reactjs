package dto

type TheMovieDBPage struct {
	PageNumber int          `json:"page"`
	Results    []TheMovieDB `json:"results"`
	TotalPages int          `json:"total_pages"`
}

type TheMovieDB struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	ReleaseDae  string  `json:"release_date,omitempty"`
	PosterPath  string  `json:"poster_path,omitempty"`
	Runtime     int     `json:"runtime,omitempty"`
	VoteAverage float32 `json:"vote_average,omitempty"`
	VoteCount   int     `json:"vote_count,omitempty"`
	Genres      []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres,omitempty"`
	Overview string `json:"overview,omitempty"`
}
