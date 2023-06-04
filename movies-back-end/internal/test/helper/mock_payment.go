package helper

import (
	"context"
	"movies-service/internal/model"
)

type MockPaymentRepository struct {
	InsertPaymentFn           func(ctx context.Context, payment *model.Payment) (*model.Payment, error)
	FindByProviderPaymentIDFn func(ctx context.Context, provider model.PaymentProvider, providerPaymentID string) (*model.Payment, error)
}

func NewMockPaymentRepository() *MockPaymentRepository {
	return &MockPaymentRepository{}
}

func (mpr *MockPaymentRepository) InsertPayment(ctx context.Context, payment *model.Payment) (*model.Payment, error) {
	if mpr.InsertPaymentFn != nil {
		return mpr.InsertPaymentFn(ctx, payment)
	}
	return nil, nil
}

func (mpr *MockPaymentRepository) FindByProviderPaymentID(ctx context.Context, provider model.PaymentProvider, providerPaymentID string) (*model.Payment, error) {
	if mpr.FindByProviderPaymentIDFn != nil {
		return mpr.FindByProviderPaymentIDFn(ctx, provider, providerPaymentID)
	}
	return nil, nil
}
