package dto

type TheMovieDBPage struct {
	PageNumber int          `json:"page"`
	Results    []TheMovieDB `json:"results"`
	TotalPages int          `json:"total_pages"`
}

type TheMovieDB struct {
	ID             int64   `json:"id"`
	Title          string  `json:"title"`
	Name           string  `json:"name"`
	ReleaseDate    string  `json:"release_date,omitempty"`
	FirstAirDate   string  `json:"first_air_date,omitempty"`
	PosterPath     string  `json:"poster_path,omitempty"`
	Runtime        int     `json:"runtime,omitempty"`
	EpisodeRuntime []int   `json:"episode_run_time,omitempty"`
	VoteAverage    float32 `json:"vote_average,omitempty"`
	VoteCount      int     `json:"vote_count,omitempty"`
	Genres         []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres,omitempty"`
	Overview string `json:"overview,omitempty"`
}
