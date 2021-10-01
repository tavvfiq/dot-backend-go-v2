package postgres

import (
	"context"
	"go-backend-v2/internal/infrastructure/database/postgres"
	"go-backend-v2/internal/model"
	"go-backend-v2/internal/user"
	"log"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
	pg postgres.PostgresInterface
}

func NewUserPostgresRepository(pg postgres.PostgresInterface) user.Repository {
	return &userPostgresRepository{pg}
}

func (r userPostgresRepository) Create(ctx context.Context, user model.User) error {
	result := r.pg.Db(ctx).Model(&model.User{}).Create(&user)
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error creating user"))
	}
	return result.Error
}
func (r userPostgresRepository) Update(ctx context.Context, user model.User) error {
	result := r.pg.Db(ctx).Model(&model.User{}).Where("id = ?", user.ID).UpdateColumns(&user)
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error updating user"))
	}
	return result.Error
}

func (r userPostgresRepository) Delete(ctx context.Context, user model.User) error {
	result := r.pg.Db(ctx).Model(&model.User{}).Delete(&user)
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error deleting user"))
	}
	return result.Error
}

func (r userPostgresRepository) FindByName(ctx context.Context, name string) (model.User, error) {
	var user model.User
	result := r.pg.Db(ctx).Model(&model.User{}).Where("name = ?", name).First(&user)
	if result.RowsAffected == 0 {
		return user, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error finding user"))
	}
	return user, result.Error
}

func (r userPostgresRepository) FindById(ctx context.Context, id int64) (model.User, error) {
	var user model.User
	result := r.pg.Db(ctx).Model(&model.User{}).Where("id = ?", id).First(&user)
	if result.RowsAffected == 0 {
		return user, errors.Wrap(gorm.ErrRecordNotFound, "user not found")
	}
	if result.Error != nil {
		log.Println(errors.Wrap(result.Error, "error finding user"))
	}
	return user, result.Error
}
