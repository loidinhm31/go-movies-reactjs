package dto

import "time"

type UserDto struct {
	ID        int       `json:"user_id,omitempty"`
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
	ID        int       `json:"role_id,omitempty"`
	RoleName  string    `json:"role_name,omitempty"`
	RoleCode  string    `json:"role_code"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Viewer struct {
	MovieId string `json:"movie_id"`
	Viewer  string `json:"viewer"`
}
