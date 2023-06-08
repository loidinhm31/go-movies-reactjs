package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/common/entity"
	"movies-service/internal/payment"
	"movies-service/pkg/pagination"
	"strings"
)

type paymentRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewPaymentRepository(cfg *config.Config, db *gorm.DB) payment.Repository {
	return &paymentRepository{cfg: cfg, db: db}
}

func (pr paymentRepository) InsertPayment(ctx context.Context, payment *entity.Payment) (*entity.Payment, error) {
	tx := pr.db.WithContext(ctx)
	if pr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Create(&payment).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (pr paymentRepository) FindPaymentByProviderPaymentID(ctx context.Context, provider entity.PaymentProvider, providerPaymentID string) (*entity.Payment, error) {
	var result *entity.Payment
	tx := pr.db.WithContext(ctx)
	if pr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("provider = ? AND provider_payment_id = ?", provider, providerPaymentID).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pr paymentRepository) FindPaymentsByTypeCodeAndRefID(ctx context.Context, typeCode string, refID uint) ([]*entity.Payment, error) {
	var results []*entity.Payment
	tx := pr.db.WithContext(ctx)
	if pr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("type_code = ? AND ref_id = ?", typeCode, refID).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (pr paymentRepository) FindPaymentByUserIDAndTypeCodeAndRefID(ctx context.Context, userID uint, typeCode string, refID uint) (*entity.Payment, error) {
	var result *entity.Payment
	tx := pr.db.WithContext(ctx)
	if pr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("user_id = ? AND type_code = ? AND ref_id = ?", userID, typeCode, refID).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pr paymentRepository) FindPaymentsByUserID(ctx context.Context, userID uint, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.CustomPayment]) (*pagination.Page[*entity.CustomPayment], error) {
	var results []*entity.CustomPayment
	var totalRows int64

	tx := pr.db.WithContext(ctx)
	if pr.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Table("payments p").
		Select("p.id as id, p.type_code as type_code, m.title as movie_title, s.name as season_name, e.name as episode_name, " +
			"p.provider as provider, p.payment_method as payment_method, p.amount as amount, p.currency as currency, p.status as status," +
			" p.created_at as created_at").
		Joins("LEFT JOIN movies m ON p.ref_id = m.id").
		Joins("LEFT JOIN episodes e ON p.ref_id = e.id").
		Joins("LEFT JOIN seasons s on e.season_id = s.id")

	if keyword != "" {
		lowerWord := fmt.Sprintf("%%%s%%", strings.ToLower(keyword))
		tx = tx.Where("LOWER(m.title) LIKE ? OR LOWER(s.name) LIKE ? OR LOWER(e.name) LIKE ?", lowerWord, lowerWord, lowerWord)
	}

	err := tx.Where("user_id = ?", userID).
		Count(&totalRows).
		Scopes(pagination.PageImplCountCriteria[*entity.CustomPayment](totalRows, pageRequest, page)).
		Find(&results).Error
	if err != nil {
		return nil, err
	}
	page.Content = results
	return page, nil
}
