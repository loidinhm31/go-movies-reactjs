package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/entity"
	"movies-service/pkg/pagination"
)

type MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) InsertPayment(ctx context.Context, payment *entity.Payment) (*entity.Payment, error) {
	args := m.Called(ctx, payment)
	return args.Get(0).(*entity.Payment), args.Error(1)
}

func (m *MockPaymentRepository) FindByProviderPaymentID(ctx context.Context, provider entity.PaymentProvider, providerPaymentID string) (*entity.Payment, error) {
	args := m.Called(ctx, provider, providerPaymentID)
	return args.Get(0).(*entity.Payment), args.Error(1)
}

func (m *MockPaymentRepository) FindByTypeCodeAndRefID(ctx context.Context, typeCode string, refID uint) (*entity.Payment, error) {
	args := m.Called(ctx, typeCode, refID)
	return args.Get(0).(*entity.Payment), args.Error(1)
}

func (m *MockPaymentRepository) FindPaymentByProviderPaymentID(ctx context.Context, provider entity.PaymentProvider, providerPaymentID string) (*entity.Payment, error) {
	args := m.Called(ctx, provider, providerPaymentID)
	return args.Get(0).(*entity.Payment), args.Error(1)
}

func (m *MockPaymentRepository) FindPaymentsByTypeCodeAndRefID(ctx context.Context, typeCode string, refID uint) ([]*entity.Payment, error) {
	args := m.Called(ctx, typeCode, refID)
	return args.Get(0).([]*entity.Payment), args.Error(1)
}

func (m *MockPaymentRepository) FindPaymentByUserIDAndTypeCodeAndRefID(ctx context.Context, userID uint, typeCode string, refID uint) (*entity.Payment, error) {
	args := m.Called(ctx, userID, typeCode, refID)
	return args.Get(0).(*entity.Payment), args.Error(1)
}

func (m *MockPaymentRepository) FindPaymentsByUserID(ctx context.Context, userID uint, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.CustomPayment]) (*pagination.Page[*entity.CustomPayment], error) {
	args := m.Called(ctx, userID, keyword, pageRequest, page)
	return args.Get(0).(*pagination.Page[*entity.CustomPayment]), args.Error(1)
}
