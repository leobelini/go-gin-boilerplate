package user

import (
	"context"
	entity "leobelini/cashly/internal/domain/user"
)

type GetByEmailUserUseCase struct {
	repo UserRepository
}

func NewGetByEmailUserUseCase(repo UserRepository) *GetByEmailUserUseCase {
	return &GetByEmailUserUseCase{repo: repo}
}

func (uc *GetByEmailUserUseCase) Execute(ctx context.Context, email string) (entity.User, error) {
	return uc.repo.GetByEmail(ctx, email)
}
