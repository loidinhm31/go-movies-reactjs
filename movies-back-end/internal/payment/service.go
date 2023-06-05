package payment

import (
	"context"
	dto2 "movies-service/internal/common/dto"
	"movies-service/internal/common/model"
)

type Service interface {
	VerifyPayment(ctx context.Context, provider model.PaymentProvider, providerPaymentID string, username string, typeCode string, refID uint) error
	GetPaymentsByUser(ctx context.Context) ([]*dto2.PaymentDto, error)
	GetPaymentsByUserAndTypeCodeAndRefID(ctx context.Context, typeCode string, refID uint) (*dto2.PaymentDto, error)
}
