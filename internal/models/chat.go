package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Chat string
}
