package adapters

import (
	"davidPardoC/rest/auth/dtos"
	"davidPardoC/rest/auth/usecases"
)

type authAdapters struct {
	uc *usecases.AuthUseCases
}

func NewAuthAdapters(uc *usecases.AuthUseCases) *authAdapters {
	return &authAdapters{uc: uc}
}

func (a *authAdapters) HandleSingup(email string) {
	a.uc.SignUpUser(email)
}

func (a *authAdapters) HandleLogin(loginDto dtos.LoginDto) string {
	return "Well done papu"
}
