package postgres

import (
	"context"
	"go-backend-v2/internal/comment"
	"go-backend-v2/internal/infrastructure/database/postgres"
	"go-backend-v2/internal/model"
	"log"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type commentPostgresRepository struct {
	pg postgres.PostgresInterface
}

func NewCommentPostgresRepository(pg postgres.PostgresInterface) comment.Repository {
	return &commentPostgresRepository{pg: pg}
}

func (r commentPostgresRepository) FindById(ctx context.Context, id int64) (model.Comment, error) {
	var comment model.Comment
	result := r.pg.Db(ctx).Model(&model.Comment{}).Where("id = ?", id).First(&comment)
	if result.RowsAffected == 0 {
		return comment, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "find comment error"))
	}
	return comment, result.Error
}

func (r commentPostgresRepository) GetByArticleId(ctx context.Context, articleId int64) ([]model.Comment, error) {
	var comments []model.Comment
	result := r.pg.Db(ctx).Model(&model.Comment{}).Where("article_id = ?", articleId).Preload("User").Find(&comments)
	if result.RowsAffected == 0 {
		return comments, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "find comment error"))
	}
	return comments, result.Error
}

func (r commentPostgresRepository) Create(ctx context.Context, comment model.Comment) error {
	result := r.pg.Db(ctx).Model(&model.Comment{}).Create(&comment)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error on create comment"))
	}
	return result.Error
}

func (r commentPostgresRepository) Update(ctx context.Context, comment model.Comment) error {
	result := r.pg.Db(ctx).Model(&model.Comment{}).Where("id = ?", comment.ID).UpdateColumns(&comment)
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error on create comment"))
	}
	return result.Error
}

func (r commentPostgresRepository) Delete(ctx context.Context, comment model.Comment) error {
	result := r.pg.Db(ctx).Model(&model.Comment{}).Delete(&comment)
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error on create comment"))
	}
	return result.Error
}
