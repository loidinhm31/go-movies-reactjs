package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Email     string
	FirstName string
	LastName  string
	IsNew     bool
	RoleID    uint
	Role      *Role `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

type Role struct {
	ID        uint `gorm:"primaryKey"`
	RoleName  string
	RoleCode  string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
