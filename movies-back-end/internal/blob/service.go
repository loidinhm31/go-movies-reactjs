package blob

import (
	"context"
	"mime/multipart"
)

type Service interface {
	UploadFile(ctx context.Context, file multipart.File, fileType string) (string, error)
	DeleteFile(ctx context.Context, fileId string, fileType string) (string, error)
}
