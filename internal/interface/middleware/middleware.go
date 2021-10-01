package middleware

import (
	"go-backend-v2/internal/domain"
	"go-backend-v2/internal/helper"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			return helper.ResponseError(c, domain.GetStatusCode(err), err.Error())
		}
		return nil
	}
}
