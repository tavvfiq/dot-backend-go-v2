package article

import (
	"context"
	"go-backend-v2/internal/model"
)

type Repository interface {
	WithTransaction(ctx context.Context, fn func(ctxWithTx context.Context) error) error
	Create(ctx context.Context, article *model.Article) error
	Update(ctx context.Context, article *model.Article) error
	Delete(ctx context.Context, article *model.Article) error
	FindById(ctx context.Context, articleId int64) (model.Article, error)
	FindByTitle(ctx context.Context, title string, limit int, offset int) ([]model.Article, error)
}
