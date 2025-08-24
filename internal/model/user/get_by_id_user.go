package user

import (
	"context"
	"errors"
	"leobelini/cashly/internal/entity"

	"gorm.io/gorm"
)

func (m *UserModel) GetByIdUser(id string, ctx context.Context) (entity.User, error) {
	var user entity.User

	err := m.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	if user.ID == "" {
		return entity.User{}, gorm.ErrRecordNotFound
	}

	return user, nil
}
