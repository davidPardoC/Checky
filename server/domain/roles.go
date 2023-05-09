package domain

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
