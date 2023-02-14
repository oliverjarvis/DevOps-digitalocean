package application

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string
	Email     string
	PW_hash   string
	Messages  []*Message
	Followers []*User `gorm:"many2many:user_followers"`
}
