package usecase

import (
	"context"
	"go-backend-v2/internal/comment"
	"go-backend-v2/internal/domain"
	"go-backend-v2/internal/model"
	"time"

	"gorm.io/gorm"
)

type commentUsecase struct {
	commentRepo comment.Repository
}

func NewCommentUsecase(commentRepo comment.Repository) comment.Usecase {
	return &commentUsecase{
		commentRepo: commentRepo,
	}
}

func (u commentUsecase) Add(ctx context.Context, req domain.AddCommentRequest) error {
	comment := model.Comment{
		UserId:    req.UserId,
		ArticleId: req.ArticleId,
		Content:   req.Content,
		CreatedAt: time.Now().UTC(),
	}
	err := u.commentRepo.Create(ctx, comment)
	if err == gorm.ErrRecordNotFound {
		return domain.ErrNotFound
	}
	if err != nil {
		return domain.ErrInternalServer
	}
	return nil
}

func (u commentUsecase) Delete(ctx context.Context, id int64) error {
	comment, err := u.commentRepo.FindById(ctx, id)
	if err == gorm.ErrRecordNotFound {
		return domain.ErrNotFound
	}
	if err != nil {
		return domain.ErrInternalServer
	}
	err = u.commentRepo.Delete(ctx, comment)
	if err != nil {
		return domain.ErrInternalServer
	}
	return nil
}
