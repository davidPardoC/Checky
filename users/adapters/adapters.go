package adapters

import "davidPardoC/rest/users/usecase"

type userAdpater struct {
	uc *usecase.UserUseCase
}

func NewUserAdapter(uc *usecase.UserUseCase) *userAdpater {
	return &userAdpater{uc: uc}
}

func (ua *userAdpater) HandleNewUser() {
	ua.uc.CreateNewUser()
}
