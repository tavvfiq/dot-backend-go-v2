package helper

import (
	"go-backend-v2/internal/domain"

	"github.com/labstack/echo/v4"
)

// ResponseWithJSON response request with JSON
func ResponseWithJSON(c echo.Context, statusCode int, payload interface{}) error {
	return c.JSON(statusCode, payload)
}

// ResponseWithError response request with error
func ResponseError(c echo.Context, statusCode int, message string) error {
	var res domain.Response
	res.Message = message
	res.Success = false
	return ResponseWithJSON(c, statusCode, res)
}

func ResponseSuccess(c echo.Context, statusCode int, message string) error {
	var res domain.Response
	res.Message = message
	res.Success = true
	return ResponseWithJSON(c, statusCode, res)
}

func ResponseSuccessWithData(c echo.Context, statusCode int, message string, data interface{}) error {
	var res domain.Response
	res.Message = message
	res.Data = data
	res.Success = true
	return ResponseWithJSON(c, statusCode, res)
}
