package user

import (
	"leobelini/cashly/internal/integration"
	"leobelini/cashly/internal/types/database"

	"gorm.io/gorm"
)

func (u *UserModel) GetUserByEmail(email string) (database.User, error) {
	return gorm.G[database.User](integration.Db).Where("email = ?", email).First(integration.DbCtx)
}
