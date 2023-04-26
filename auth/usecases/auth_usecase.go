package usecases

import (
	"davidPardoC/rest/domain"
	"davidPardoC/rest/users/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCases struct {
	userRepo repository.UserRepository
}

func NewAuthUseCases(userRepo repository.UserRepository) *AuthUseCases {
	return &AuthUseCases{userRepo: userRepo}
}

func (useCase *AuthUseCases) SignUpUser(user domain.User) (int64, error) {
	user.Password, _ = hashPassword(user.Password)
	return useCase.userRepo.InsertNewUser(user)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
