package adapters

import "davidPardoC/rest/auth/usecases"

type authAdapters struct {
	uc *usecases.AuthUseCases
}

func NewAuthAdapters(uc *usecases.AuthUseCases) *authAdapters {
	return &authAdapters{uc: uc}
}

func (a *authAdapters) HandleSingup(email string) {
	a.uc.SignUpUser(email)
}
