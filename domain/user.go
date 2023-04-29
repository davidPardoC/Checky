package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
