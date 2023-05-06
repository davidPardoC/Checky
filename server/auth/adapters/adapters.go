package adapters

import (
	"davidPardoC/rest/auth/dtos"
	"davidPardoC/rest/auth/usecases"
	"davidPardoC/rest/common"
	"davidPardoC/rest/domain"
)

type authAdapters struct {
	useCase *usecases.AuthUseCases
}

func NewAuthAdapters(uc *usecases.AuthUseCases) *authAdapters {
	uc.CreateInitialAdminUser()
	return &authAdapters{useCase: uc}
}

func (adapter *authAdapters) HandleSingup(singUpDto dtos.SignupDto) (common.SuccesMessage, error) {
	user := domain.User{Name: singUpDto.Name, LastName: singUpDto.LastName, Password: singUpDto.Password, Email: singUpDto.Email}
	rowsAffected, err := adapter.useCase.SignUpUser(user)

	if rowsAffected > 0 {
		return common.CreateSuccesCreatedMessage("User"), nil
	}

	return common.CreateSuccesCreatedMessage("User not"), err

}

func (adapter *authAdapters) HandleLogin(loginDto dtos.LoginDto) (*dtos.TokenResponse, *common.CustomError) {
	token, err := adapter.useCase.LoginUser(loginDto)
	return token, err
}
