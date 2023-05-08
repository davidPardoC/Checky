package repository

import (
	"davidPardoC/rest/common"
	"davidPardoC/rest/domain"
)

type UserRepository interface {
	InsertNewUser(domain.User) (int64, error)
	GetUserByEmail(email string) (*domain.User, *common.CustomError)
	GetAllUsers() (*[]domain.User, *common.CustomError)
}
