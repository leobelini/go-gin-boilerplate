package db

import (
	"context"
	"errors"
	domainUser "leobelini/cashly/internal/domain/user"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (r *GormUserRepository) Save(ctx context.Context, user domainUser.User) error {
	model := UserModel{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		CreatedAt:       user.CreatedAt,
		Token:           user.Token,
		Password:        user.Password,
		AccountVerified: user.AccountVerified,
	}
	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *GormUserRepository) GetByEmail(ctx context.Context, email string) (domainUser.User, error) {
	var model UserModel

	err := r.db.WithContext(ctx).Where("email = ?", email).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domainUser.User{}, nil
		}
		return domainUser.User{}, err
	}

	if model.ID == "" {
		return domainUser.User{}, gorm.ErrRecordNotFound
	}

	return domainUser.User{ID: model.ID, Name: model.Name, Email: model.Email, CreatedAt: model.CreatedAt, AccountVerified: model.AccountVerified, Token: model.Token, Password: model.Password}, nil
}
