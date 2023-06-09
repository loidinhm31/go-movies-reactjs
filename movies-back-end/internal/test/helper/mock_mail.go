package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/model"
)

type MockMailService struct {
	mock.Mock
}

func (m *MockMailService) SendMessage(ctx context.Context, mail *model.MailData) error {
	args := m.Called(ctx, mail)
	return args.Error(0)
}
