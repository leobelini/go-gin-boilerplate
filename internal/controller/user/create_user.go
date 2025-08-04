package user

import (
	"context"
	"leobelini/cashly/internal/entity"
	"leobelini/cashly/internal/model"
	"leobelini/cashly/internal/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (c *UserController) CreateUser(name, email, password string, ctx context.Context) error {

	return c.model.Transaction(func(tx *model.Model) error {
		userExists, err := c.model.User.GetByEmailUser(email, ctx)
		if err != nil {
			return err
		}

		if userExists != (entity.User{}) {
			return utils.CreateAppError("USER_ALREADY_EXISTS", false)
		}

		bytesPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := entity.User{
			ID:              uuid.New().String(),
			Name:            name,
			Email:           email,
			CreatedAt:       time.Now(),
			Password:        string(bytesPassword),
			AccountVerified: false,
			Token:           uuid.New().String(),
		}

		if err := c.model.User.CreateUser(user, ctx); err != nil {
			return err
		}

		userRegistered, err := c.model.User.GetByEmailUser(email, ctx)
		if err != nil {
			return err
		}

		if userRegistered == (entity.User{}) {
			return utils.CreateAppError("USER_NOT_FOUND", false)
		}

		c.job.SendConfirmationEmailRegister.SendConfirmationEmailRegister(ctx, email, name, userRegistered.Token)

		return nil
	})

}
