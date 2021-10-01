package usecase

import (
	"context"
	"go-backend-v2/internal/article"
	"go-backend-v2/internal/comment"
	"go-backend-v2/internal/domain"
	"go-backend-v2/internal/model"
	"time"

	"gorm.io/gorm"
)

type articleUsecase struct {
	articleRepo article.Repository
	commentRepo comment.Repository
}

func NewArticleUsecase(articleRepo article.Repository, commentRepo comment.Repository) article.Usecase {
	return &articleUsecase{articleRepo, commentRepo}
}

func (u articleUsecase) CreateNewArticle(ctx context.Context, req domain.NewArticleRequestData) error {
	article := model.Article{
		Title:     req.Title,
		Subtitle:  req.Subtitle,
		AuthorId:  req.AuthorId,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}
	err := u.articleRepo.Create(ctx, &article)
	if err == gorm.ErrRecordNotFound {
		return domain.ErrNotFound
	}
	if err != nil {
		return domain.ErrInternalServer
	}
	return nil
}

func (u articleUsecase) GetArticleDetail(ctx context.Context, id int64) (domain.ArticleDetailData, error) {
	var articleDetailData domain.ArticleDetailData
	err := u.articleRepo.WithTransaction(ctx, func(ctxWithTx context.Context) error {
		article, err := u.articleRepo.FindById(ctxWithTx, id)
		if err == gorm.ErrRecordNotFound {
			return domain.ErrNotFound
		}
		if err != nil {
			return err
		}
		comments, err := u.commentRepo.GetByArticleId(ctxWithTx, article.ID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		var articleComments []domain.ArticleComment
		for i := 0; i < len(comments); i++ {
			articleComments = append(articleComments, domain.ArticleComment{
				Id:        comments[i].ID,
				UserName:  comments[i].User.Name,
				UserId:    comments[i].UserId,
				Content:   comments[i].Content,
				CreatedAt: comments[i].CreatedAt,
			})
		}
		articleDetailData = domain.ArticleDetailData{
			Id:         article.ID,
			AuthorName: article.Author.Name,
			AuthorId:   article.AuthorId,
			Title:      article.Title,
			Subtitle:   article.Subtitle,
			Content:    article.Content,
			Comments:   articleComments,
		}
		return nil
	})
	return articleDetailData, err
}

func (u articleUsecase) GetArticle(ctx context.Context, req domain.ArticleRequestData) (domain.ArticleResponseData, error) {
	offset := (req.Page - 1) * req.Limit
	nilArticleResponseData := domain.ArticleResponseData{
		Articles:   nil,
		Pagination: domain.Pagination{},
	}
	articles, err := u.articleRepo.FindByTitle(ctx, req.Title, req.Limit, offset)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nilArticleResponseData, domain.ErrInternalServer
	}
	var articlesData []domain.ArticleData
	for i := 0; i < len(articles); i++ {
		articlesData = append(articlesData, domain.ArticleData{
			Id:         articles[i].ID,
			Title:      articles[i].Title,
			Subtitle:   articles[i].Subtitle,
			AuthorName: articles[i].Author.Name,
			CreatedAt:  articles[i].CreatedAt,
			UpdatedAt:  articles[i].UpdatedAt,
		})
	}
	last := req.Page - 1
	if last < 0 {
		last = 0
	}
	next := req.Page + 1
	if len(articlesData) == 0 {
		next = 0
	}
	return domain.ArticleResponseData{
		Articles: articlesData,
		Pagination: domain.Pagination{
			Last:    last,
			Current: req.Page,
			Next:    next,
		},
	}, nil
}

func (u articleUsecase) UpdateArticle(ctx context.Context, id int64, req domain.UpdateArticleRequestData) error {
	article, err := u.articleRepo.FindById(ctx, id)
	if err == gorm.ErrRecordNotFound {
		return domain.ErrNotFound
	}
	if req.Title != nil {
		article.Title = *req.Title
	}
	if req.Subtitle != nil {
		article.Subtitle = *req.Subtitle
	}
	if req.Content != nil {
		article.Content = *req.Content
	}
	if req.AuthorId != nil {
		article.AuthorId = *req.AuthorId
	}
	article.UpdatedAt = time.Now().UTC()
	err = u.articleRepo.Update(ctx, &article)
	if err != nil {
		return domain.ErrInternalServer
	}
	return nil
}

func (u articleUsecase) DeleteArticle(ctx context.Context, id int64) error {
	article, err := u.articleRepo.FindById(ctx, id)
	if err == gorm.ErrRecordNotFound {
		return domain.ErrNotFound
	}
	if err != nil {
		return domain.ErrInternalServer
	}
	err = u.articleRepo.Delete(ctx, &article)
	if err != nil {
		return domain.ErrInternalServer
	}
	return nil
}
