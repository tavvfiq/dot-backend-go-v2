package postgres

import (
	"context"
	"go-backend-v2/internal/article"
	"go-backend-v2/internal/infrastructure/database/postgres"
	"go-backend-v2/internal/model"
	"log"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type articlePostgresRepository struct {
	pg postgres.PostgresInterface
}

func NewArticlePostgresRepository(pg postgres.PostgresInterface) article.Repository {
	return &articlePostgresRepository{pg}
}

func (r articlePostgresRepository) WithTransaction(ctx context.Context, fn func(ctxWithTx context.Context) error) error {
	tx := r.pg.Db(ctx).Begin()
	ctxWithTx := context.WithValue(ctx, "txCtx", tx)
	err := fn(ctxWithTx)
	if err != nil {
		if err := tx.Rollback().Error; err != nil {
			if err != nil {
				return errors.Wrap(err, "error on rollback")
			}
		}
		return err
	}
	if err := tx.Commit().Error; err != nil {
		if err != nil {
			return errors.Wrap(err, "error on commit")
		}
		return err
	}
	return nil
}

func (r articlePostgresRepository) Create(ctx context.Context, article *model.Article) error {
	result := r.pg.Db(ctx).Model(&model.Article{}).Create(&article)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error on create article"))
	}
	return result.Error
}

func (r articlePostgresRepository) Update(ctx context.Context, article *model.Article) error {
	result := r.pg.Db(ctx).Model(&model.Article{}).Where("id = ?", article.ID).UpdateColumns(&article)
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error on update article"))
	}
	return result.Error
}

func (r articlePostgresRepository) Delete(ctx context.Context, article *model.Article) error {
	result := r.pg.Db(ctx).Model(&model.Article{}).Delete(&article)
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error on delete article"))
	}
	return result.Error
}

func (r articlePostgresRepository) FindById(ctx context.Context, articleId int64) (model.Article, error) {
	var article model.Article
	result := r.pg.Db(ctx).Model(&model.Article{}).Where("id = ?", articleId).Preload("Author").First(&article)
	if result.RowsAffected == 0 {
		return article, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error on find article"))
		return article, result.Error
	}
	return article, nil
}

func (r articlePostgresRepository) FindByTitle(ctx context.Context, title string, limit int, offset int) ([]model.Article, error) {
	var articles []model.Article
	// must concat "%" to both start and end of title. otherwise, wont find anything
	_title := "%" + title + "%"
	result := r.pg.Db(ctx).Model(&model.Article{}).Where("title like ?", _title).Limit(limit).Offset(offset).Preload("Author").Find(&articles)
	if result.RowsAffected == 0 {
		return articles, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error on find article"))
		return articles, result.Error
	}
	return articles, nil
}
