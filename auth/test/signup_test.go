package auth_test

import (
	"bytes"
	"davidPardoC/rest/domain"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignupRouteOnBadBody(t *testing.T) {

	SetupEmptyUserDatabaseTests()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", basePath+"/auth/signup", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestSignupRouteOnCorrectBody(t *testing.T) {

	SetupEmptyUserDatabaseTests()

	w := httptest.NewRecorder()

	expectedBody := gin.H{
		"Message": "User created successfully",
	}

	body := []byte(`{
		"name":"david",
		"lastname":"pardo",
		"email":"pardodavid10@gmail.com",
		"password":"mypassword"
	}`)

	req, _ := http.NewRequest("POST", basePath+"/auth/signup", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Grab the value & whether or not it exists
	value, exists := response["Message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedBody["Message"], value)
	assert.Equal(t, 200, w.Code)

}

func TestDuplicatedEmail(t *testing.T) {
	SetupEmptyUserDatabaseTests()
	database.Create(&domain.User{Name: "david", LastName: "pardo", Email: "pardodavid10@gmail.com", Password: "mypassword"})

	w := httptest.NewRecorder()

	body := []byte(`{
		"name":"david",
		"lastname":"pardo",
		"email":"pardodavid10@gmail.com",
		"password":"mypassword"
	}`)

	req, _ := http.NewRequest("POST", basePath+"/auth/signup", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)

	assert.Equal(t, 409, w.Code)
}
