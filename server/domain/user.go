package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"not null"`
	LastName string `json:"lastname" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`
}

type NoPasswordUser struct {
	ID       uint   `json:"id"`
	Email    string `json:"email" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"not null"`
	LastName string `json:"lastname" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`
}

func (user User) ToNoPasswordUser() NoPasswordUser {
	noPasswordUser := NoPasswordUser{ID: user.ID, Email: user.Email, Name: user.Name, LastName: user.LastName, Role: user.Role}
	return noPasswordUser
}
