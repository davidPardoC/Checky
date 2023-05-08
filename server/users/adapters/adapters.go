package adapters

import (
	"davidPardoC/rest/common"
	"davidPardoC/rest/domain"
	"davidPardoC/rest/users/usecases"
)

type userAdpater struct {
	usecase *usecases.UserUseCase
}

func NewUserAdapter(usecase *usecases.UserUseCase) *userAdpater {
	return &userAdpater{usecase: usecase}
}

func (adapter *userAdpater) HandleNewUser() {
}

func (adapter *userAdpater) HandleGetUserById(id string) {
}

func (adapter *userAdpater) GetAllUsers() (*[]domain.NoPasswordUser, *common.CustomError) {
	return adapter.usecase.GetAllUsers()
}
