package entity

import (
	"time"
)

type User struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password,omitempty"`
	AccountVerified bool      `json:"accountVerified,omitempty"`
	Token           string    `json:"token,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
}

func (u User) TableName() string {
	return "users"
}
