package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/errors"
	"movies-service/internal/model"
	"movies-service/internal/search"
	"movies-service/pkg/pagination"
	"strings"
)

const (
	Id          = "id"
	TypeCode    = "type_code"
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

func (sr *searchRepository) SearchMovie(ctx context.Context, searchParams *model.SearchParams) (*pagination.Page[*model.Movie], error) {
	page := &pagination.Page[*model.Movie]{}
	var movies []*model.Movie
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
		case TypeCode:
			if f.TypeValue.Values[0] != "" {
				tx.Where("type_code = ?", f.TypeValue.Values[0])
			}
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
			if len(f.TypeValue.Values) > 0 {
				subQuery := sr.db.Table("movies").
					Select("movies.id").
					Joins("JOIN movies_genres mg ON movies.id = mg.movie_id").
					Joins("JOIN genres g ON g.id = mg.genre_id")

				var orBuild = sr.db
				for _, g := range f.TypeValue.Values {
					genreSplit := strings.Split(g, "-")
					orBuild = orBuild.Or("g.name = ? AND g.type_code = ?", genreSplit[0], genreSplit[1])
				}

				subQuery = subQuery.Where(orBuild)
				tx.Where("id IN (?)", subQuery)
			}
			break
		}
	}

	err := tx.Count(&totalRows).
		Scopes(pagination.PageImplCountCriteria[*model.Movie](totalRows, searchParams.Page, page)).
		Preload("Genres").
		Find(&movies).Error
	if err != nil {
		fmt.Println(err)
	}
	page.Content = movies
	page.TotalElements = totalRows
	return page, nil
}

func (sr *searchRepository) buildLikeQuery(tx *gorm.DB, field, operator string, def model.TypeValue) error {
	if len(def.Values) > 0 && def.Type != model.DATE {
		var orBuild = sr.db
		for _, val := range def.Values {
			val = fmt.Sprintf("%%%s%%", strings.ToLower(val))
			orBuild = orBuild.Or("LOWER("+field+") LIKE ?", val)
		}

		if strings.EqualFold(operator, model.AND) {
			tx.Where(orBuild)
		} else if strings.EqualFold(operator, model.OR) {
			tx.Or(orBuild)
		}
		return nil
	}
	return errors.ErrInvalidInput
}

func (sr *searchRepository) buildEqualQuery(tx *gorm.DB, field, operator string, def model.TypeValue) error {
	if len(def.Values) > 0 && def.Type != model.DATE {
		var orBuild = sr.db

		if def.Type == model.NUMBER {
			for _, val := range def.Values {
				orBuild = orBuild.Or(field+" = ?", strings.ToLower(val))
			}
		} else {
			for _, val := range def.Values {
				orBuild = orBuild.Or("LOWER("+field+") = ?", strings.ToLower(val))
			}
		}

		if strings.EqualFold(operator, model.AND) {
			tx.Where(orBuild)
		} else if strings.EqualFold(operator, model.OR) {
			tx.Or(orBuild)
		}
		return nil
	}
	return errors.ErrInvalidInput
}

func buildDateQuery(tx *gorm.DB, field, operator string, def model.TypeValue) error {
	if len(def.Values) == 2 && def.Type == model.DATE {
		if def.Values[0] == "" || def.Values[1] == "" {
			return errors.ErrInvalidInput
		}

		if strings.EqualFold(operator, model.AND) {
			tx = tx.Where(field+" BETWEEN ? AND ?", def.Values[0], def.Values[1])
		} else if strings.EqualFold(operator, model.OR) {
			tx = tx.Or(field+" BETWEEN ? AND ?", def.Values[0], def.Values[1])
		}
		return nil
	}
	return errors.ErrInvalidInput
}
