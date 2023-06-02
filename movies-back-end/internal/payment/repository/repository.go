package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/model"
	"movies-service/internal/payment"
)

type paymentRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewPaymentRepository(cfg *config.Config, db *gorm.DB) payment.Repository {
	return &paymentRepository{cfg: cfg, db: db}
}

func (pr paymentRepository) InsertPayment(ctx context.Context, payment *model.Payment) (*model.Payment, error) {
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

func (pr paymentRepository) FindByProviderPaymentID(ctx context.Context, provider model.PaymentProvider, providerPaymentID string) (*model.Payment, error) {
	var result *model.Payment
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
