package integration

import (
	"context"
	"mime/multipart"
)

type Service interface {
	UploadVideo(ctx context.Context, file multipart.File) (string, error)
	DeleteVideo(ctx context.Context, fileId string) (string, error)
}
