package dto

import "time"

type SeasonDto struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	AirDate     time.Time     `json:"air_date"`
	Description string        `json:"description"`
	MovieID     uint          `json:"movie_id"`
	EpisodeDtos []*EpisodeDto `json:"episodes"`
}

type EpisodeDto struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	AirDate   time.Time `json:"air_date"`
	Runtime   uint      `json:"runtime"`
	VideoPath string    `json:"video_path"`
	SeasonID  uint      `json:"season_id"`
}
