package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	ActiveJob   bool   `json:"active_job"`
	Users       []User `gorm:"many2many:project_users" json:"users"`
}
