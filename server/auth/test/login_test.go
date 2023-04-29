package auth_test

import (
	"bytes"
	"davidPardoC/rest/auth/usecases"
	"davidPardoC/rest/domain"
	"davidPardoC/rest/users/repository/postgres"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginNotFoudEmail(t *testing.T) {
	SetupEmptyUserDatabaseTests()

	w := httptest.NewRecorder()

	endpoint := GetFinalEndpoint("/login")

	fmt.Println(endpoint)

	body := []byte(`{
		"email":"pardodavid10@gmail.com",
		"password":"mypassword"		
	}`)

	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))

	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)

}

func TestLoginUserBadBody(t *testing.T) {
	SetupEmptyUserDatabaseTests()

	w := httptest.NewRecorder()

	endpoint := GetFinalEndpoint("/login")

	fmt.Println(endpoint)

	body := []byte(`{
		"email":"pardodavid10@gmail.com",
		"bad_key":"mypassword"		
	}`)

	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

}

func TestLoginUser(t *testing.T) {
	SetupEmptyUserDatabaseTests()

	userRepository := postgres.NewUserPostgresRepository(database)
	authUseCases := usecases.NewAuthUseCases(userRepository)

	user := domain.User{Email: "pardodavid10@gmail.com", Password: "mypassword"}

	authUseCases.SignUpUser(user)

	w := httptest.NewRecorder()

	endpoint := GetFinalEndpoint("/login")

	fmt.Println(endpoint)

	body := []byte(`{
		"email":"pardodavid10@gmail.com",
		"password":"mypassword"		
	}`)

	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestLoginUserBadCredentials(t *testing.T) {
	SetupEmptyUserDatabaseTests()

	userRepository := postgres.NewUserPostgresRepository(database)
	authUseCases := usecases.NewAuthUseCases(userRepository)

	user := domain.User{Email: "pardodavid10@gmail.com", Password: "mypassword"}

	authUseCases.SignUpUser(user)

	w := httptest.NewRecorder()

	endpoint := GetFinalEndpoint("/login")

	fmt.Println(endpoint)

	body := []byte(`{
		"email":"pardodavid10@gmail.com",
		"password":"bad_password"		
	}`)

	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
