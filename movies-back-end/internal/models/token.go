package models

type UserToken struct {
	Claims Claims `json:"Claims"`
}

type Claims struct {
	Username string `json:"preferred_username"`
}
