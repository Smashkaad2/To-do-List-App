package models

import (
	"time"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID        uint
	TaskName  string
	Status	  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Task) TableName() string {
	return "tasks"
}