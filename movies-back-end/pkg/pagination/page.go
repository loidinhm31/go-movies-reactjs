package pagination

import (
	"gorm.io/gorm"
	"math"
)

type Page[K any] struct {
	PageSize   int   `json:"size,omitempty"`
	PageNumber int   `json:"page,omitempty"`
	Sort       Sort  `json:"sort,omitempty"`
	TotalData  int64 `json:"total_data"`
	TotalPages int   `json:"total_pages"`
	Data       []K   `json:"data"`
}

func PageImpl[K any](value []K, pageRequest *PageRequest, page *Page[K], db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	page.TotalData = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageRequest.PageSize)))
	page.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pageRequest.GetOffset()).Limit(pageRequest.GetLimit()).Order(pageRequest.GetSort())
	}
}
