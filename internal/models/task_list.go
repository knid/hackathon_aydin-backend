package models

import "gorm.io/gorm"

type TaskList struct {
	gorm.Model
	List[] Task
	SortBy string
}
