package entity

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	TypeCode    string
	ReleaseDate time.Time
	Runtime     uint
	MpaaRating  string
	Description string
	ImageUrl    sql.NullString  `gorm:"type:varchar(255), default:null"`
	VideoPath   sql.NullString  `gorm:"type:varchar(255), default:null"`
	Price       sql.NullFloat64 `gorm:"type:float, default:null"`
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
