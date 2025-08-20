package user

import "gorm.io/gorm"

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(dbGorm *gorm.DB) *UserModel {
	return &UserModel{db: dbGorm}
}
