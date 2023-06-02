package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/collection"
	"movies-service/internal/model"
	"movies-service/pkg/pagination"
	"strings"
)

type collectionRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewCollectionRepository(cfg *config.Config, db *gorm.DB) collection.Repository {
	return &collectionRepository{cfg: cfg, db: db}
}

func (fr collectionRepository) InsertCollection(ctx context.Context, collection *model.Collection) error {
	tx := fr.db.WithContext(ctx)
	if fr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Create(&collection).Error
	if err != nil {
		return err
	}
	return nil
}

func (fr collectionRepository) FindCollectionsByUsernameAndType(ctx context.Context, username string, movieType string, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*model.CollectionDetail]) (*pagination.Page[*model.CollectionDetail], error) {
	var results []*model.CollectionDetail
	var totalRows int64

	tx := fr.db.WithContext(ctx)
	if fr.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.WithContext(ctx).Table("collections c").
		Joins("JOIN movies m ON m.id = c.movie_id").
		Joins("JOIN payments p ON p.id = c.payment_id").
		Select("m.title, m.description, m.release_date, m.image_url, p.amount, c.created_at").
		Where("c.username = ? AND m.type_code = ?", username, movieType)

	if keyword != "" {
		lowerWord := fmt.Sprintf("%%%s%%", strings.ToLower(keyword))
		tx = tx.Where("LOWER(m.title) LIKE ? OR LOWER(m.description) = ?", lowerWord, lowerWord)
	}

	err := tx.
		Scopes(pagination.PageImplCountCriteria[*model.CollectionDetail](totalRows, pageRequest, page)).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	page.Content = results
	return page, nil
}

func (fr collectionRepository) FindCollectionByUsernameAndMovieID(ctx context.Context, username string, movieID uint) (*model.Collection, error) {
	var result *model.Collection
	tx := fr.db.WithContext(ctx)
	if fr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("username = ? AND movie_id = ?", username, movieID).
		Find(&result).Error
	if err != nil {
		return result, nil
	}
	return result, nil
}

func (fr collectionRepository) FindCollectionByPaymentID(ctx context.Context, paymentID uint) (*model.Collection, error) {
	var result *model.Collection
	tx := fr.db.WithContext(ctx)
	if fr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("payment_id = ?", paymentID).
		Find(&result).Error
	if err != nil {
		return result, nil
	}
	return result, nil
}

func (fr collectionRepository) FindCollectionsByMovieID(ctx context.Context, movieID uint) ([]*model.Collection, error) {
	var results []*model.Collection
	tx := fr.db.WithContext(ctx)
	if fr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("movie_id = ?", movieID).
		Find(&results).Error
	if err != nil {
		return results, nil
	}
	return results, nil
}
