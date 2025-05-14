package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title string `json:"name"`
	Description string `json:"description"`
	EstimatedTime uint32 `json:"est_time"`
	ProjectID uint
	Project Project `json:"project"`
	Status int8 `json:"status"`
	Priority int32 `json:"priority"`
}
