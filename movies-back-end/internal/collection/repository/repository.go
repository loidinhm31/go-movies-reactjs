package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/collection"
	"movies-service/internal/common/entity"
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

func (cr *collectionRepository) InsertCollection(ctx context.Context, collection *entity.Collection) error {
	tx := cr.db.WithContext(ctx)
	if cr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Create(&collection).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *collectionRepository) FindCollectionsByUserIDAndType(ctx context.Context, userID uint, movieType string, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.CollectionDetail]) (*pagination.Page[*entity.CollectionDetail], error) {
	var results []*entity.CollectionDetail
	var totalRows int64

	tx := cr.db.WithContext(ctx)
	if cr.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.WithContext(ctx).Table("collections c")
	if movieType == "MOVIE" {
		tx = tx.Joins("JOIN movies m ON m.id = c.movie_id AND c.type_code = ?", movieType).
			Joins("LEFT JOIN payments p ON m.id = p.ref_id AND p.type_code = ?", movieType).
			Select("c.id, m.id as movie_id, m.title, m.description, m.release_date, m.image_url, p.amount, c.created_at").
			Where("c.user_id = ?", userID)
	} else if movieType == "TV" {
		tx = tx.Joins("JOIN episodes e ON e.id = c.episode_id AND c.type_code = ?", movieType).
			Joins("JOIN seasons s ON s.id = e.season_id").
			Joins("JOIN movies m ON m.id = s.movie_id").
			Joins("LEFT JOIN payments p ON e.id = p.ref_id AND p.type_code = ?", movieType).
			Select("c.id, m.id as movie_id, e.id as episode_id, m.title, m.description, s.name as season_name, e.name as episode_name, e.air_date as release_date, m.image_url, p.amount, c.created_at").
			Where("c.user_id = ?", userID)
	}

	if keyword != "" {
		lowerWord := fmt.Sprintf("%%%s%%", strings.ToLower(keyword))
		tx = tx.Where("LOWER(m.title) LIKE ? OR LOWER(m.description) LIKE ?", lowerWord, lowerWord)
	}

	err := tx.
		Scopes(pagination.PageImplCountCriteria[*entity.CollectionDetail](totalRows, pageRequest, page)).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	page.Content = results
	return page, nil
}

func (cr *collectionRepository) FindCollectionByUserIDAndMovieID(ctx context.Context, userID uint, movieID uint) (*entity.Collection, error) {
	var result *entity.Collection
	tx := cr.db.WithContext(ctx)
	if cr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("user_id = ? AND movie_id = ?", userID, movieID).
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (cr *collectionRepository) FindCollectionByUserIDAndEpisodeID(ctx context.Context, userID uint, episodeID uint) (*entity.Collection, error) {
	var result *entity.Collection
	tx := cr.db.WithContext(ctx)
	if cr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("user_id = ? AND episode_id = ?", userID, episodeID).
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (cr *collectionRepository) FindCollectionByPaymentID(ctx context.Context, paymentID uint) (*entity.Collection, error) {
	var result *entity.Collection
	tx := cr.db.WithContext(ctx)
	if cr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("payment_id = ?", paymentID).
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (cr *collectionRepository) FindCollectionsByMovieID(ctx context.Context, movieID uint) ([]*entity.Collection, error) {
	var results []*entity.Collection
	tx := cr.db.WithContext(ctx)
	if cr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("movie_id = ?", movieID).
		Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (cr *collectionRepository) FindCollectionByMovieID(ctx context.Context, episodeID uint) (*entity.Collection, error) {
	var results *entity.Collection
	tx := cr.db.WithContext(ctx)
	if cr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("episode_id = ?", episodeID).
		Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (cr *collectionRepository) FindCollectionsByEpisodeID(ctx context.Context, episodeID uint) ([]*entity.Collection, error) {
	var results []*entity.Collection
	tx := cr.db.WithContext(ctx)
	if cr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("episode_id = ?", episodeID).
		Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (cr *collectionRepository) FindCollectionByEpisodeID(ctx context.Context, episodeID uint) (*entity.Collection, error) {
	var results *entity.Collection
	tx := cr.db.WithContext(ctx)
	if cr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("episode_id = ?", episodeID).
		Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (cr *collectionRepository) FindCollectionsByID(ctx context.Context, id uint) (*entity.Collection, error) {
	var results *entity.Collection
	tx := cr.db.WithContext(ctx)
	if cr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("id = ?", id).
		Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (cr *collectionRepository) DeleteCollectionByTypeCodeAndMovieID(ctx context.Context, typeCode string, movieID uint) error {
	tx := cr.db.WithContext(ctx)
	if cr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("type_code = ? AND movie_id = ?", typeCode, movieID).Delete(&entity.Collection{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *collectionRepository) DeleteCollectionByTypeCodeAndEpisodeID(ctx context.Context, typeCode string, episodeID uint) error {
	tx := cr.db.WithContext(ctx)
	if cr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("type_code = ? AND episode_id = ?", typeCode, episodeID).Delete(&entity.Collection{}).Error
	if err != nil {
		return err
	}
	return nil
}
