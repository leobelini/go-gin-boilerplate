package infra

import (
	db "leobelini/cashly/internal/infra/db"

	"gorm.io/gorm"
)

type RepositoryGorm struct {
	User *db.GormUserRepository
}

func NewRepositoryGorm(dbGorm *gorm.DB) *RepositoryGorm {
	return &RepositoryGorm{User: db.NewGormUserRepository(dbGorm)}
}
