package user

import (
	"context"
	"leobelini/cashly/internal/entity"
	"leobelini/cashly/internal/utils"
)

func (c *UserController) GetUserByEmail(email string, ctx context.Context) (entity.User, error) {

	user, err := c.app.Model.User.GetByEmailUser(email, ctx)
	if err != nil {
		return entity.User{}, err
	}

	if user.ID == "" {
		return entity.User{}, utils.CreateAppError("USER_NOT_FOUND", false)
	}

	return user, nil

}
