package payment

import (
	"context"
	"movies-service/internal/common/model"
)

type Repository interface {
	InsertPayment(ctx context.Context, payment *model.Payment) (*model.Payment, error)
	FindPaymentByProviderPaymentID(ctx context.Context, provider model.PaymentProvider, providerPaymentID string) (*model.Payment, error)
	FindPaymentByTypeCodeAndRefID(ctx context.Context, typeCode string, refID uint) (*model.Payment, error)
	FindPaymentByUserIDAndTypeCodeAndRefID(ctx context.Context, userID uint, typeCode string, refID uint) (*model.Payment, error)
	FindPaymentsByUserID(ctx context.Context, userID uint) ([]*model.Payment, error)
}
