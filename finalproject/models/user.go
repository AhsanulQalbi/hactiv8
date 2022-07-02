package models

import (
	"finalproject/helpers"

	"gorm.io/gorm"
)

type User struct {
	ID         int    `gorm:"primary key" json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	Age        int    `json:"age,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(ctx *gorm.DB) (err error) {
	u.Password = helpers.PasswordHashing(u.Password)
	err = nil
	return
}
