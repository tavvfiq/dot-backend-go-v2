package usecase

import (
	"context"
	"go-backend-v2/internal/domain"
	"go-backend-v2/internal/user"
	"log"

	"gorm.io/gorm"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return &userUsecase{userRepo}
}

func (u userUsecase) Create(ctx context.Context, req domain.RegisterRequestData) error {
	user, err := u.userRepo.FindByName(ctx, req.Name)
	if err == gorm.ErrRecordNotFound {
		// create new user
		user.Name = req.Name
		err = u.userRepo.Create(ctx, user)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	if err != nil {
		return domain.ErrInternalServer
	}
	return domain.ErrConflict
}
