package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `gorm:"unique;"`
	Password string `json:"-"`
}
