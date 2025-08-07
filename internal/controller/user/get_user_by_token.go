package user

import (
	"context"
	"leobelini/cashly/internal/entity"
	"leobelini/cashly/internal/utils"
)

func (c *UserController) GetUserByToken(token string, ctx context.Context) (entity.User, error) {

	user, err := c.model.User.GetByTokenUser(token, ctx)
	if err != nil {
		return entity.User{}, err
	}

	if user.ID == "" {
		return entity.User{}, utils.CreateAppError("USER_NOT_FOUND", false)
	}

	return user, nil

}
