package usecase

import "davidPardoC/rest/users/repository"

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (uc *UserUseCase) CreateNewUser() {
	uc.repository.InsertNewUser()
}
