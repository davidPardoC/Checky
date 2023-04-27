package auth_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
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

	assert.Equal(t, 200, w.Code)

}

func TestLoginUserBadBody(t *testing.T) {
	SetupEmptyUserDatabaseTests()

	w := httptest.NewRecorder()

	endpoint := GetFinalEndpoint("/login")

	fmt.Println(endpoint)

	body := []byte(`{
		"email":"pardodavid10@gmail.com",
		"password1":"mypassword"		
	}`)

	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

}
