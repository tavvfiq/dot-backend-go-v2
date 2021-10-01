package user

import (
	"context"
	"go-backend-v2/internal/model"
)

type Repository interface {
	Create(ctx context.Context, user model.User) error
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, user model.User) error
	FindByName(ctx context.Context, name string) (model.User, error)
	FindById(ctx context.Context, id int64) (model.User, error)
}
