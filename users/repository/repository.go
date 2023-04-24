package repository

import "davidPardoC/rest/users/domain"

type UserRepository interface {
	InsertNewUser() string
	GetUserByEmail(email string) domain.User
	GetAllUsers() []domain.User
}
