package payment

import (
	"context"
	"movies-service/internal/model"
)

type Service interface {
	VerifyPayment(ctx context.Context, provider model.PaymentProvider, providerPaymentID string, username string, typeCode string, refID uint) error
}
