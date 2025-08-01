package infra

import (
	"context"
	entity "leobelini/cashly/internal/domain/user"

	"gorm.io/gorm"
)

type UserRepository interface {
	WithTransaction(ctx context.Context, fn func(UserRepository) error) error
	Save(ctx context.Context, user entity.User) error
	GetByEmail(ctx context.Context, email string) (entity.User, error)
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}
