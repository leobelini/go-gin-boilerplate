package user

import (
	"context"
	"leobelini/cashly/internal/entity"
)

func (m *UserModel) CreateUser(user entity.User, ctx context.Context) error {
	return m.db.WithContext(ctx).Create(&user).Error
}
