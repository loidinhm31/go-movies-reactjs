package model

import "time"

type View struct {
	ID       uint64 `gorm:"primaryKey"`
	ViewedBy string
	ViewedAt time.Time
	MovieId  uint
}
