package models

import (
	"time"
)

type User struct {
	ID        int `gorm:"primary_key"`
	Username  string
	Email     string
	FirstName string
	LastName  string
	RoleID    int
	Role      Role `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	ID        int `gorm:"primary_key"`
	RoleName  string
	RoleCode  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
