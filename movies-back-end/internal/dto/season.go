package dto

import "time"

type SeasonDto struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	AirDate     time.Time     `json:"air_date"`
	Description string        `json:"description"`
	MovieID     int           `json:"movie_id"`
	EpisodeDtos []*EpisodeDto `json:"episodes"`
}

type EpisodeDto struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	AirDate   time.Time `json:"air_date"`
	Runtime   int       `json:"runtime"`
	VideoPath string    `json:"video_path"`
	SeasonID  int       `json:"season_id"`
}
