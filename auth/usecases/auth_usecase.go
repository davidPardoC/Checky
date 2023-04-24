package usecases

import (
	"davidPardoC/rest/users/repository"
)

type AuthUseCases struct {
	userRepo *repository.UserRepository
}

func NewAuthUseCases(userRepo repository.UserRepository) *AuthUseCases {
	return &AuthUseCases{}
}

func (uc *AuthUseCases) SignUpUser(email string) {

}
