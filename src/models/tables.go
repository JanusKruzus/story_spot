package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey"`
	Name     string `gorm:"size:64"`
	Email    string `gorm:"size:64"`
	Password string `gorm:"size:255"`
}
