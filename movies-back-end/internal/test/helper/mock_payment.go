package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/model"
)

type MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) InsertPayment(ctx context.Context, payment *model.Payment) (*model.Payment, error) {
	args := m.Called(ctx, payment)
	return args.Get(0).(*model.Payment), args.Error(1)
}

func (m *MockPaymentRepository) FindByProviderPaymentID(ctx context.Context, provider model.PaymentProvider, providerPaymentID string) (*model.Payment, error) {
	args := m.Called(ctx, provider, providerPaymentID)
	return args.Get(0).(*model.Payment), args.Error(1)
}

func (m *MockPaymentRepository) FindByTypeCodeAndRefID(ctx context.Context, typeCode string, refID uint) (*model.Payment, error) {
	args := m.Called(ctx, typeCode, refID)
	return args.Get(0).(*model.Payment), args.Error(1)
}
