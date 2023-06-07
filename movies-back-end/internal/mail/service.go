package mail

import (
	"context"
	"movies-service/internal/common/model"
)

type Service interface {
	SendMessage(ctx context.Context, mail *model.MailData) error
}
