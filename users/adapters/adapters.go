package adapters

import "davidPardoC/rest/users/usecase"

type userAdpater struct{}

func NewUserAdapter() *userAdpater {
	return &userAdpater{}
}

func (ua *userAdpater) HandleNewUser() {
	userUseCases := usecase.NewUserUseCase()
	userUseCases.CreateNewUser()
}
