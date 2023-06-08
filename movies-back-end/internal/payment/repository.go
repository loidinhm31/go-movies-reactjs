package payment

import (
	"context"
	"movies-service/internal/common/entity"
	"movies-service/pkg/pagination"
)

type Repository interface {
	InsertPayment(ctx context.Context, payment *entity.Payment) (*entity.Payment, error)
	FindPaymentByProviderPaymentID(ctx context.Context, provider entity.PaymentProvider, providerPaymentID string) (*entity.Payment, error)
	FindPaymentsByTypeCodeAndRefID(ctx context.Context, typeCode string, refID uint) ([]*entity.Payment, error)
	FindPaymentByUserIDAndTypeCodeAndRefID(ctx context.Context, userID uint, typeCode string, refID uint) (*entity.Payment, error)
	FindPaymentsByUserID(ctx context.Context, userID uint, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.CustomPayment]) (*pagination.Page[*entity.CustomPayment], error)
}
