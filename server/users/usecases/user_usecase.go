package usecases

import (
	"davidPardoC/rest/common"
	"davidPardoC/rest/domain"
	"davidPardoC/rest/users/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (usecase *UserUseCase) GetAllUsers() (*[]domain.NoPasswordUser, *common.CustomError) {
	users, err := usecase.repository.GetAllUsers()
	mappedUsers := []domain.NoPasswordUser{}
	for _, user := range *users {
		mappedUsers = append(mappedUsers, user.ToNoPasswordUser())
	}
	return &mappedUsers, err
}
