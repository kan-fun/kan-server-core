package model

import "time"

type Log struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time

	UserID uint
	User   User
	Topic  string `gorm:"type:varchar(100);not null"`
}
