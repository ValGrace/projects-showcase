package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}
