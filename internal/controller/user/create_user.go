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

	return c.app.Model.Transaction(func(tx *model.Model) error {
		userExists, err := c.app.Model.User.GetByEmailUser(email, ctx)
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

		token := uuid.New().String()

		user := entity.User{
			ID:              uuid.New().String(),
			Name:            name,
			Email:           email,
			CreatedAt:       time.Now(),
			Password:        string(bytesPassword),
			AccountVerified: false,
			Token:           &token,
		}

		if err := c.app.Model.User.CreateUser(user, ctx); err != nil {
			return err
		}

		userRegistered, err := c.app.Model.User.GetByEmailUser(email, ctx)
		if err != nil {
			return err
		}

		if userRegistered == (entity.User{}) {
			return utils.CreateAppError("USER_NOT_FOUND", false)
		}

		if err := c.app.Job.SendConfirmationEmailRegister.AddQueue(ctx, email, name, *userRegistered.Token); err != nil {
			return err
		}
		return nil
	})

}
