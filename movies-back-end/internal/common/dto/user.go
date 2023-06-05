package dto

import "time"

type UserDto struct {
	ID        uint      `json:"user_id,omitempty"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	IsNew     bool      `json:"is_new"`
	Role      RoleDto   `json:"role"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type RoleDto struct {
	ID        uint      `json:"role_id,omitempty"`
	RoleName  string    `json:"role_name,omitempty"`
	RoleCode  string    `json:"role_code"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Viewer struct {
	MovieID uint   `json:"movie_id"`
	Viewer  string `json:"viewer"`
}

type CollectionDto struct {
	UserID    uint   `json:"user_id,omitempty"`
	MovieID   uint   `json:"movie_id,omitempty"`
	EpisodeID uint   `json:"episode_id,omitempty"`
	TypeCode  string `json:"type_code"`
}

type CollectionDetailDto struct {
	UserID      string    `json:"user_id,omitempty"`
	MovieID     uint      `json:"id,omitempty"`
	EpisodeID   uint      `json:"episode_id,omitempty"`
	TypeCode    string    `json:"type_code,omitempty"`
	Title       string    `json:"title,omitempty"`
	SeasonName  string    `json:"season_name,omitempty"`
	EpisodeName string    `json:"episode_name"`
	ReleaseDate time.Time `json:"release_date"`
	ImageUrl    string    `json:"image_url,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
