package domain

import (
	"errors"
	"net/http"
)

var (
	ErrBadParam       = errors.New("bad param given")
	ErrInternalServer = errors.New("internal server error")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrConflict       = errors.New("already exist")
	ErrNotFound       = errors.New("not found")
)

func GetStatusCode(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrBadParam:
		return http.StatusBadRequest
	case ErrInternalServer:
		return http.StatusInternalServerError
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
