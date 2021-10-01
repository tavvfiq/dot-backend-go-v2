package cache

import (
	"context"
	"fmt"
	"go-backend-v2/internal/article"
	"go-backend-v2/internal/infrastructure/database/redis"
	"go-backend-v2/internal/model"
)

type articleRepositoryWithCache struct {
	articleRepo article.Repository
	cache       redis.CacheProvider
}

func NewArticleRepositoryWithCache(cache redis.CacheProvider, articleRepo article.Repository) article.Repository {
	return &articleRepositoryWithCache{
		articleRepo: articleRepo,
		cache:       cache,
	}
}

func (r articleRepositoryWithCache) WithTransaction(ctx context.Context, fn func(ctxWithTx context.Context) error) error {
	err := r.articleRepo.WithTransaction(ctx, fn)
	return err
}

func (r articleRepositoryWithCache) Create(ctx context.Context, article *model.Article) error {
	err := r.articleRepo.Create(ctx, article)
	if err != nil {
		return err
	}
	// find cache
	var _article model.Article
	key := fmt.Sprintf("article-%d", article.ID)
	r.cache.Get(ctx, key, &_article)
	if _article.ID != 0 {
		// delete cache if available
		r.cache.Delete(ctx, key)
	}
	return nil
}

func (r articleRepositoryWithCache) Update(ctx context.Context, article *model.Article) error {
	err := r.articleRepo.Update(ctx, article)
	if err != nil {
		return err
	}
	// find cache
	var _article model.Article
	key := fmt.Sprintf("article-%d", article.ID)
	r.cache.Get(ctx, key, &_article)
	if _article.ID != 0 {
		// delete cache if available
		r.cache.Delete(ctx, key)
	}
	return nil
}

func (r articleRepositoryWithCache) Delete(ctx context.Context, article *model.Article) error {
	err := r.articleRepo.Delete(ctx, article)
	if err != nil {
		return err
	}
	// find cache
	var _article model.Article
	key := fmt.Sprintf("article-%d", article.ID)
	r.cache.Get(ctx, key, &_article)
	if _article.ID != 0 {
		// delete cache if available
		r.cache.Delete(ctx, key)
	}
	return nil
}

func (r articleRepositoryWithCache) FindById(ctx context.Context, articleId int64) (model.Article, error) {
	// find cache
	var _article model.Article
	key := fmt.Sprintf("article-%d", articleId)
	err := r.cache.Get(ctx, key, &_article)
	if err == nil {
		return _article, nil
	}
	_article, err = r.articleRepo.FindById(ctx, articleId)
	if err != nil {
		return _article, err
	}
	r.cache.Store(ctx, key, &_article)
	return _article, nil
}

func (r articleRepositoryWithCache) FindByTitle(ctx context.Context, title string, limit int, offset int) ([]model.Article, error) {
	articles, err := r.articleRepo.FindByTitle(ctx, title, limit, offset)
	if err != nil {
		return articles, err
	}
	return articles, nil
}
