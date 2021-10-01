package http

import (
	"go-backend-v2/internal/domain"
	"go-backend-v2/internal/helper"
	"go-backend-v2/internal/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHttpHandler struct {
	userUsecase user.Usecase
}

func NewUserHttpHandler(e *echo.Group, userUsecase user.Usecase) {
	handler := &userHttpHandler{userUsecase}
	e.POST("/user", handler.CreateUser)
}

func (h userHttpHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	var req domain.RegisterRequestData
	err := c.Bind(&req)
	if err != nil {
		return domain.ErrBadParam
	}
	err = h.userUsecase.Create(ctx, req)
	if err != nil {
		return err
	}
	return helper.ResponseSuccess(c, http.StatusCreated, "user created")
}
