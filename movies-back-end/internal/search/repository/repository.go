package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/model"
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
	Genres      = "genre"
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
	tx = tx.Table("movie")

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
			if len(f.TypeValue.Values) > 0 {
				tx.Where("id IN (SELECT m.id FROM movie m "+
					"JOIN movies_genres mg on m.id = mg.movie_id "+
					"JOIN genre g on g.id = mg.genre_id WHERE genre IN ?)", f.TypeValue.Values)
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
	return errors.New("invalid input")
}

func (sr *searchRepository) buildEqualQuery(tx *gorm.DB, field, operator string, def model.TypeValue) error {
	if len(def.Values) > 0 && def.Type != model.DATE {
		var orBuild = sr.db
		for _, val := range def.Values {
			orBuild = orBuild.Or(field+" = ?", strings.ToLower(val))
		}

		if strings.EqualFold(operator, model.AND) {
			tx.Where(orBuild)
		} else if strings.EqualFold(operator, model.OR) {
			tx.Or(orBuild)
		}
		return nil
	}
	return errors.New("invalid input")
}

func buildDateQuery(tx *gorm.DB, field, operator string, def model.TypeValue) error {
	if len(def.Values) == 2 && def.Type == model.DATE {
		if strings.EqualFold(operator, model.AND) {
			tx = tx.Where(field+" BETWEEN ? AND ?", def.Values[0], def.Values[1])
		} else if strings.EqualFold(operator, model.OR) {
			tx = tx.Or(field+" BETWEEN ? AND ?", def.Values[0], def.Values[1])
		}
		return nil
	}
	return errors.New("invalid input")
}
