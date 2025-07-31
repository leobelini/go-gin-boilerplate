package user

import (
	"context"
	entity "leobelini/cashly/internal/domain/user"
)

type UserRepository interface {
	Save(ctx context.Context, user entity.User) error
	GetByEmail(ctx context.Context, email string) (entity.User, error)
}

type UserUseCase struct {
	CreateUser *CreateUserUseCase
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
	return &UserUseCase{CreateUser: NewCreateUserUseCase(repo)}
}
