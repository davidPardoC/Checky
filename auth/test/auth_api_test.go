package test

import (
	"bytes"
	"davidPardoC/rest/common"
	"davidPardoC/rest/config"
	"davidPardoC/rest/domain"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var basePath string = "/api/v1"

var database, _ = common.GetTestDatabase()

var router = config.SetupRouter(database)

func TestSignupRouteOnBadBody(t *testing.T) {

	config.MakeMigrations(database)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", basePath+"/auth/signup", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestSignupRouteOnCorrectBody(t *testing.T) {
	database.Migrator().DropTable(domain.User{})
	config.MakeMigrations(database)

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
