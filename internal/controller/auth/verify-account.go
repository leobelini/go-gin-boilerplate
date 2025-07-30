package auth

import (
	"errors"
	_userModel "leobelini/cashly/internal/model/user"
	"leobelini/cashly/internal/utils"

	"gorm.io/gorm"
)

func (a *AuthController) VerifyAccount(token string) error {
	userModel := &_userModel.UserModel{}

	user, err := userModel.GetUserByToken(token)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.CreateAppError("INVALID_TOKEN", false)
		}
		return err
	}

	if user.AccountVerified {
		return utils.CreateAppError("INVALID_TOKEN", false)
	}
	// tx := integration.Db.Session(&gorm.Session{})
	return nil
}
