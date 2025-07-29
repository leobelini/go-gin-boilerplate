package user

import (
	"errors"
	_userModel "leobelini/cashly/internal/model/user"
	"leobelini/cashly/internal/types/database"
	"leobelini/cashly/internal/utils"

	"gorm.io/gorm"
)

func (u *UserController) CreateUser(user *database.User) (database.User, error) {
	userModel := &_userModel.UserModel{}

	userExists, err := userModel.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return database.User{}, err
	}

	if userExists.ID != "" {
		return database.User{}, utils.CreateAppError("USER_ALREADY_EXISTS", false)
	}

	createdUser, err := userModel.InsertUser(*user)
	if err != nil {
		return database.User{}, err
	}

	return createdUser, nil
}
