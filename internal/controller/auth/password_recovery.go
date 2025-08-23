package auth

import (
	"context"
	"leobelini/cashly/internal/model"
	"leobelini/cashly/internal/utils"

	"github.com/google/uuid"
)

func (c *AuthController) PasswordRecovery(email string, ctx context.Context) error {

	return c.app.Model.Transaction(func(tx *model.Model) error {
		user, err := c.app.Model.User.GetByEmailUser(email, ctx)
		if err != nil {
			return err
		}

		if user.ID == "" || !user.AccountVerified {
			return utils.CreateAppError("USER_NOT_FOUND", false)
		}

		tokenPassword := uuid.New().String()
		user.TokenPassword = &tokenPassword

		if err := c.app.Model.User.UpdateUser(user, ctx); err != nil {
			return err
		}

		if err := c.app.Job.RecoveryPassword(ctx, user.Email, user.Name, *user.TokenPassword); err != nil {
			return err // rollback será disparado
		}

		return nil
	})
}
