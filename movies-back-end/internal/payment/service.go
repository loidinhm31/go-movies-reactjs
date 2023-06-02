package payment

import (
	"context"
	"movies-service/internal/model"
)

type Service interface {
	VerifyPayment(ctx context.Context, provider model.PaymentProvider, providerPaymentID string, username string, movieID uint) error
}
