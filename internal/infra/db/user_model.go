package db

import "time"

type UserModel struct {
	ID              string `gorm:"primaryKey"`
	Name            string
	Email           string `gorm:"unique"`
	Password        string
	AccountVerified bool
	Token           string
	CreatedAt       time.Time
}

func (u *UserModel) TableName() string {
	return "users"
}
