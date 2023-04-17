package usecase

import "davidPardoC/rest/users/repository/postgres"

type UserUseCase struct {
	UserRepository postgres.UserPostgresRepository
}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{}
}

func (uc *UserUseCase) CreateNewUser() {
	uc.UserRepository.InsertNewUser()
}
