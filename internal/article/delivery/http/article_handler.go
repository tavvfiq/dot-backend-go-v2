package http

import (
	"go-backend-v2/internal/article"
	"go-backend-v2/internal/domain"
	"go-backend-v2/internal/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type articleHttpHandler struct {
	articleUsecase article.Usecase
}

func NewArticleHttpHandler(e *echo.Group, articleUsecase article.Usecase) {
	handler := &articleHttpHandler{articleUsecase}
	e.POST("/article", handler.CreateNewArticle)
	e.GET("/article", handler.GetAllArticle)
	e.GET("/article/detail/:id", handler.GetDetailArticle)
	e.PATCH("/article/:id", handler.UpdateArticle)
	e.DELETE("/article/:id", handler.DeleteArticle)
}

func (h articleHttpHandler) CreateNewArticle(c echo.Context) error {
	ctx := c.Request().Context()
	var req domain.NewArticleRequestData
	err := c.Bind(&req)
	if err != nil {
		return domain.ErrBadParam
	}
	err = h.articleUsecase.CreateNewArticle(ctx, req)
	if err != nil {
		return err
	}
	return helper.ResponseSuccess(c, http.StatusCreated, "article created")
}

func (h articleHttpHandler) GetAllArticle(c echo.Context) error {
	ctx := c.Request().Context()
	title := c.QueryParam("title")
	limit, err := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	if err != nil {
		limit = 10
	}
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	if err != nil {
		page = 1
	}
	req := domain.ArticleRequestData{
		Title: title,
		Limit: int(limit),
		Page:  int(page),
	}
	articles, err := h.articleUsecase.GetArticle(ctx, req)
	if err != nil {
		return err
	}
	return helper.ResponseSuccessWithData(c, http.StatusOK, "request success", articles)
}

func (h articleHttpHandler) GetDetailArticle(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return domain.ErrBadParam
	}
	articleDetail, err := h.articleUsecase.GetArticleDetail(ctx, id)
	if err != nil {
		return err
	}
	return helper.ResponseSuccessWithData(c, http.StatusOK, "request success", articleDetail)
}

func (h articleHttpHandler) UpdateArticle(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return domain.ErrBadParam
	}
	var req domain.UpdateArticleRequestData
	err = c.Bind(&req)
	if err != nil {
		return domain.ErrBadParam
	}
	err = h.articleUsecase.UpdateArticle(ctx, id, req)
	if err != nil {
		return err
	}
	return helper.ResponseSuccess(c, http.StatusOK, "article updated")
}

func (h articleHttpHandler) DeleteArticle(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return domain.ErrBadParam
	}
	err = h.articleUsecase.DeleteArticle(ctx, id)
	if err != nil {
		return err
	}
	return helper.ResponseSuccess(c, http.StatusOK, "article deleted")
}
