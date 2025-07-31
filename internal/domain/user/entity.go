package user

import (
	"time"
)

type User struct {
	ID              string
	Name            string
	Email           string
	Password        string
	AccountVerified bool
	Token           string
	CreatedAt       time.Time
}
