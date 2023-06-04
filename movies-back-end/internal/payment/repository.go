package payment

import (
	"context"
	"movies-service/internal/model"
)

type Repository interface {
	InsertPayment(ctx context.Context, payment *model.Payment) (*model.Payment, error)
	FindByProviderPaymentID(ctx context.Context, provider model.PaymentProvider, providerPaymentID string) (*model.Payment, error)
	FindByTypeCodeAndRefID(ctx context.Context, typeCode string, refID uint) (*model.Payment, error)
}
