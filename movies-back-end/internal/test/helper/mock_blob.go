package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
)

type MockBlobService struct {
	mock.Mock
}

func (m *MockBlobService) UploadFile(ctx context.Context, file multipart.File, fileType string) (string, error) {
	args := m.Called(ctx, file, fileType)
	return args.String(0), args.Error(1)
}

func (m *MockBlobService) DeleteFile(ctx context.Context, file string, fileType string) (string, error) {
	args := m.Called(ctx, file, fileType)
	return args.String(0), args.Error(1)
}
