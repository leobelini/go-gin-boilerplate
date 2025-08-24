package auth

import (
	"context"
	"leobelini/cashly/internal/utils"
)

func (c *AuthController) SignIn(email, password string, ctx context.Context) (string, string, error) {

	user, err := c.app.Model.User.GetByEmailUser(email, ctx)
	if err != nil {
		return "", "", utils.CreateAppError("INVALID_CREDENTIALS", false)
	}

	if user.ID == "" || !user.AccountVerified {
		return "", "", utils.CreateAppError("INVALID_CREDENTIALS", false)
	}

	if err := utils.CheckPasswordHash(password, user.Password); err != nil {
		return "", "", utils.CreateAppError("INVALID_CREDENTIALS", false)
	}

	token, refreshToken, err := utils.GenerateJWT(user.ID, c.app.Env.App.JWTSecret)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}
