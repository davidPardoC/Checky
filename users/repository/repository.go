package repository

import "davidPardoC/rest/domain"

type UserRepository interface {
	InsertNewUser(domain.User) (int64, error)
}
