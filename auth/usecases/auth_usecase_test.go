package usecases_test

import (
	"davidPardoC/rest/auth/usecases"
	"davidPardoC/rest/domain"
	"testing"
)

type mockUserRepository struct {
}

func (repo mockUserRepository) GetAllUsers() []domain.User {
	return []domain.User{}
}

func (repo mockUserRepository) GetUserByEmail(email string) domain.User {
	return domain.User{}
}

func (repo mockUserRepository) InsertNewUser(domain.User) (int64, error) {
	return 1, nil
}

func TestAddNewUser(t *testing.T) {
	useCases := usecases.NewAuthUseCases(mockUserRepository{})

	newUser := domain.User{}

	got, _ := useCases.SignUpUser(newUser)

	if got < 0 {
		t.Errorf("got %d, wanted %d", got, 1)
	}

}
