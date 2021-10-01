package http

import (
	"go-backend-v2/internal/comment"
	"go-backend-v2/internal/domain"
	"go-backend-v2/internal/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type commentHtppHandler struct {
	commentUsecase comment.Usecase
}

func NewCommentHttpHandler(e *echo.Group, commentUsecase comment.Usecase) {
	handler := &commentHtppHandler{commentUsecase}
	e.POST("/comment", handler.AddComment)
	e.DELETE("/comment/:id", handler.DeleteComment)
}

func (h commentHtppHandler) AddComment(c echo.Context) error {
	ctx := c.Request().Context()
	var req domain.AddCommentRequest
	err := c.Bind(&req)
	if err != nil {
		return domain.ErrBadParam
	}
	err = h.commentUsecase.Add(ctx, req)
	if err != nil {
		return err
	}
	return helper.ResponseSuccess(c, http.StatusCreated, "comment added")
}

func (h commentHtppHandler) DeleteComment(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return domain.ErrBadParam
	}
	err = h.commentUsecase.Delete(ctx, id)
	if err != nil {
		return err
	}
	return helper.ResponseSuccess(c, http.StatusCreated, "comment deleted")
}
