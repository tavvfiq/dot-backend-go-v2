package article

import (
	"context"
	"go-backend-v2/internal/domain"
)

type Usecase interface {
	CreateNewArticle(ctx context.Context, req domain.NewArticleRequestData) error
	GetArticleDetail(ctx context.Context, id int64) (domain.ArticleDetailData, error)
	GetArticle(ctx context.Context, req domain.ArticleRequestData) (domain.ArticleResponseData, error)
	UpdateArticle(ctx context.Context, id int64, req domain.UpdateArticleRequestData) error
	DeleteArticle(ctx context.Context, id int64) error
}
