package comment

import (
	"context"
	"go-backend-v2/internal/model"
)

type Repository interface {
	GetByArticleId(ctx context.Context, articleId int64) ([]model.Comment, error)
	Create(ctx context.Context, comment model.Comment) error
	Update(ctx context.Context, comment model.Comment) error
	Delete(ctx context.Context, comment model.Comment) error
	FindById(ctx context.Context, id int64) (model.Comment, error)
}
