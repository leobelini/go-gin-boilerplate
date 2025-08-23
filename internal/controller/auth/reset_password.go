package auth

import (
	"context"
	"leobelini/cashly/internal/model"
	"leobelini/cashly/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

func (c *AuthController) ResetPassword(password, token string, ctx context.Context) error {

	return c.app.Model.Transaction(func(tx *model.Model) error {
		user, err := c.app.Model.User.GetByTokenPasswordUser(token, ctx)
		if err != nil {
			return err
		}

		if user.ID == "" || !user.AccountVerified || user.TokenPassword == nil || *user.TokenPassword != token {
			return utils.CreateAppError("USER_NOT_FOUND", false)
		}

		user.TokenPassword = nil

		// Criptografa a senha
		bytesPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user.Password = string(bytesPassword)

		if err := c.app.Model.User.UpdateUser(user, ctx); err != nil {
			return err
		}

		return nil
	})
}
