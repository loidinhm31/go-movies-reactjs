package model

import "time"

type View struct {
	ID       int `gorm:"primary_key"`
	ViewedBy string
	ViewedAt time.Time
	MovieId  int
}
