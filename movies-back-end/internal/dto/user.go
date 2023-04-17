package dto

import "time"

type UserDto struct {
	UserId    string    `json:"user_id,omitempty"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// UserWithToken Find user query
type UserWithToken struct {
	User  *UserDto `json:"user"`
	Token string   `json:"token"`
}
