package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	ID          int `gorm:"primary_key"`
	Title       string
	TypeCode    string
	ReleaseDate time.Time
	Runtime     int
	MpaaRating  string
	Description string
	ImageUrl    sql.NullString `gorm:"type:varchar(255), default:null"`
	VideoPath   sql.NullString `gorm:"type:varchar(255), default:null"`
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
	Genres      []*Genre `gorm:"many2many:movies_genres;"`
}

func (m Movie) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Where("movie_id = ?", m.ID).Delete(&View{})
	return
}
