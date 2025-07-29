package user

import (
	"leobelini/cashly/internal/integration"
	"leobelini/cashly/internal/types/database"

	"gorm.io/gorm"
)

func (u *UserModel) InsertUser(user database.User) (database.User, error) {
	err := gorm.G[database.User](integration.Db).Create(integration.DbCtx, &user)
	return user, err
}
