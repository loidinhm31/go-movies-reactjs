package model

import (
	"time"
)

type User struct {
	ID        int `gorm:"primary_key"`
	Username  string
	Email     string
	FirstName string
	LastName  string
	IsNew     bool
	RoleID    int
	Role      *Role `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

type Role struct {
	ID        int `gorm:"primary_key"`
	RoleName  string
	RoleCode  string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
