package usecase

import (
	"leobelini/cashly/internal/infra"
	user "leobelini/cashly/internal/usecase/user"
)

type UseCase struct {
	User *user.UserUseCase
}

func NewUseCase(gormRepository *infra.RepositoryGorm) *UseCase {
	return &UseCase{User: user.NewUserUseCase(gormRepository.User)}
}
