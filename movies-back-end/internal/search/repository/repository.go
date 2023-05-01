package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/models"
	"movies-service/internal/search"
	"movies-service/pkg/pagination"
	"strings"
)

const (
	Id          = "id"
	Title       = "title"
	Description = "description"
	Runtime     = "runtime"
	MpaaRating  = "mpaa_rating"
	ReleaseDate = "release_date"
	Genres      = "genres"
)

// searchRepository is the type for our graphql operations
type searchRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewSearchRepository(cfg *config.Config, db *gorm.DB) search.Repository {
	return &searchRepository{cfg: cfg, db: db}
}

func (sr *searchRepository) Search(ctx context.Context, searchParams *models.SearchParams) (*pagination.Page[*models.Movie], error) {
	page := &pagination.Page[*models.Movie]{}
	var movies []*models.Movie
	var totalRows int64

	tx := sr.db.WithContext(ctx)
	if sr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	tx = tx.Table("movies")

	for _, f := range searchParams.Filters {
		switch f.Field {
		case Id:
			break
		case Title:
			err := sr.buildLikeQuery(tx, Title, f.Operator, f.TypeValue)
			if err != nil {
				return nil, err
			}
			break
		case Description:
			err := sr.buildLikeQuery(tx, Description, f.Operator, f.TypeValue)
			if err != nil {
				return nil, err
			}
			break
		case MpaaRating:
			err := sr.buildEqualQuery(tx, MpaaRating, f.Operator, f.TypeValue)
			if err != nil {
				return nil, err
			}
			break
		case Runtime:
			err := sr.buildEqualQuery(tx, Runtime, f.Operator, f.TypeValue)
			if err != nil {
				return nil, err
			}
			break
		case ReleaseDate:
			err := buildDateQuery(tx, ReleaseDate, f.Operator, f.TypeValue)
			if err != nil {
				return nil, err
			}
			break
		case Genres:
			break
		}
	}

	err := tx.Count(&totalRows).
		Scopes(pagination.PageImplCountCriteria[*models.Movie](totalRows, searchParams.Page, page)).
		Preload("Genres").
		Find(&movies).Error
	if err != nil {
		fmt.Println(err)
	}
	page.Data = movies
	page.TotalElements = totalRows
	return page, nil
}

func (sr *searchRepository) buildLikeQuery(tx *gorm.DB, field, operator string, def models.TypeValue) error {
	if len(def.Values) > 0 && def.Type != models.DATE {
		var orBuild = sr.db
		for _, val := range def.Values {
			tempVal := fmt.Sprintf("%%%s%%", val)
			orBuild = orBuild.Or("LOWER("+field+") LIKE ?", tempVal)
		}

		if strings.EqualFold(operator, models.AND) {
			tx.Where(orBuild)
		} else if strings.EqualFold(operator, models.OR) {
			tx.Or(orBuild)
		}
		return nil
	}
	return errors.New("invalid input")
}

func (sr *searchRepository) buildEqualQuery(tx *gorm.DB, field, operator string, def models.TypeValue) error {
	if len(def.Values) > 0 && def.Type != models.DATE {
		var orBuild = sr.db
		for _, val := range def.Values {
			orBuild = orBuild.Or(field+" = ?", val)
		}

		if strings.EqualFold(operator, models.AND) {
			tx.Where(orBuild)
		} else if strings.EqualFold(operator, models.OR) {
			tx.Or(orBuild)
		}
		return nil
	}
	return errors.New("invalid input")
}

func buildDateQuery(tx *gorm.DB, field, operator string, def models.TypeValue) error {
	if len(def.Values) == 2 && def.Type == models.DATE {
		if strings.EqualFold(operator, models.AND) {
			tx = tx.Where(field+" BETWEEN ? AND ?", def.Values[0], def.Values[1])
		} else if strings.EqualFold(operator, models.OR) {
			tx = tx.Or(field+" BETWEEN ? AND ?", def.Values[0], def.Values[1])
		}
		return nil
	}
	return errors.New("invalid input")
}
