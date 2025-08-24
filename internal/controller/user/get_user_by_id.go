package user

import (
	"context"
	"leobelini/cashly/internal/entity"
	"leobelini/cashly/internal/utils"
)

func (c *UserController) GetUserById(id string, ctx context.Context) (entity.User, error) {

	user, err := c.app.Model.User.GetByIdUser(id, ctx)
	if err != nil {
		return entity.User{}, err
	}

	if user.ID == "" {
		return entity.User{}, utils.CreateAppError("USER_NOT_FOUND", false)
	}

	return user, nil

}
