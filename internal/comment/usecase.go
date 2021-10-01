package comment

import (
	"context"
	"go-backend-v2/internal/domain"
)

type Usecase interface {
	Add(ctx context.Context, req domain.AddCommentRequest) error
	Delete(ctx context.Context, id int64) error
}
