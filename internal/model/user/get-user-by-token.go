package user

import (
	"leobelini/cashly/internal/integration"
	"leobelini/cashly/internal/types/database"

	"gorm.io/gorm"
)

func (u *UserModel) GetUserByToken(token string) (database.User, error) {
	return gorm.G[database.User](integration.Db).Where("token = ?", token).First(integration.DbCtx)
}
