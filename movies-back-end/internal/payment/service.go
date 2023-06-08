package payment

import (
	"context"
	dto2 "movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
	"movies-service/pkg/pagination"
)

type Service interface {
	VerifyPayment(ctx context.Context, provider entity.PaymentProvider, providerPaymentID string, username string, typeCode string, refID uint) error
	GetPaymentsByUser(ctx context.Context, keyword string, pageRequest *pagination.PageRequest) (*pagination.Page[*dto2.CustomPaymentDto], error)
	GetPaymentsByUserAndTypeCodeAndRefID(ctx context.Context, typeCode string, refID uint) (*dto2.PaymentDto, error)
}
