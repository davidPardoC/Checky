package usecases

import "davidPardoC/rest/users/repository/postgres"

type AuthUseCases struct {
	userRepo *postgres.UserPostgresRepository
}

func NewAuthUseCases() *AuthUseCases {
	return &AuthUseCases{}
}

func (uc *AuthUseCases) SignUpUser(email string) {
	uc.userRepo.GetUserByEmail(email)
}
