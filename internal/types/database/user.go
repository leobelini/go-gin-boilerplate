package database

type User struct {
	ID              string `gorm:"primaryKey"`
	Name            string
	Email           string `gorm:"unique"`
	Password        string
	AccountVerified bool
	Token           string
}

func (u *User) TableName() string {
	return "users"
}
