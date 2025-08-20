package user

import (
	"context"
	"leobelini/cashly/internal/entity"
)

func (m *UserModel) UpdateUser(user entity.User, ctx context.Context) error {
	return m.db.WithContext(ctx).Save(&user).Error
}
