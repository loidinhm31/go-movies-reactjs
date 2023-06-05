package dto

type TheMovieDBPage struct {
	PageNumber uint         `json:"page"`
	Results    []TheMovieDB `json:"results"`
	TotalPages uint         `json:"total_pages"`
}

type TheMovieDB struct {
	ID             int64   `json:"id"`
	Title          string  `json:"title"`
	Name           string  `json:"name"`
	ReleaseDate    string  `json:"release_date,omitempty"`
	FirstAirDate   string  `json:"first_air_date,omitempty"`
	PosterPath     string  `json:"poster_path,omitempty"`
	Runtime        uint    `json:"runtime,omitempty"`
	EpisodeRuntime []uint  `json:"episode_run_time,omitempty"`
	VoteAverage    float32 `json:"vote_average,omitempty"`
	VoteCount      uint    `json:"vote_count,omitempty"`
	Genres         []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"genres,omitempty"`
	Overview string `json:"overview,omitempty"`
}
