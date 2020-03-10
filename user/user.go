package user

import "time"

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time

	Email string
	Password string

	AccessKey string
	SecretKey string
}