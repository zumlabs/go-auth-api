package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:100"`
	Email    string `gorm:"uniqueIndex;size:100"`
	Password string `gorm:"size:255"`
}
