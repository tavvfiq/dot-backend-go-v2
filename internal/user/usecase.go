package user

import (
	"context"
	"go-backend-v2/internal/domain"
)

type Usecase interface {
	Create(ctx context.Context, req domain.RegisterRequestData) error
}
