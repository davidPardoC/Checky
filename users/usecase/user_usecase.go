package usecase

import "davidPardoC/rest/users/repository/postgres"

type userUseCase struct {
	UserRepository postgres.UserPostgresRepository
}

func NewUserUseCase() *userUseCase {
	return &userUseCase{}
}

func (uc *userUseCase) CreateNewUser() {
	uc.UserRepository.InsertNewUser()
}
