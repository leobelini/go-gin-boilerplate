package user

import (
	"context"
	"time"

	domainUser "leobelini/cashly/internal/domain/user"
	"leobelini/cashly/internal/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	repo UserRepository
}

func NewCreateUserUseCase(repo UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, name, email, password string) error {

	userExists, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if userExists != (domainUser.User{}) {
		return utils.CreateAppError("USER_ALREADY_EXISTS", false)
	}

	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := domainUser.User{
		ID:              uuid.New().String(),
		Name:            name,
		Email:           email,
		CreatedAt:       time.Now(),
		Password:        string(bytesPassword),
		AccountVerified: false,
		Token:           uuid.New().String(),
	}

	if err := uc.repo.Save(ctx, user); err != nil {
		return err
	}

	userRegistered, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if userRegistered == (domainUser.User{}) {
		return utils.CreateAppError("USER_NOT_FOUND", false)
	}
	// uc.job.SendConfirmationEmailRegister.SendConfirmationEmailRegister(ctx, userRegistered.Email, userRegistered.Name, userRegistered.Token)

	return nil
}
