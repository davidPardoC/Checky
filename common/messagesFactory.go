package common

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Resource string

const (
	UserResource       Resource = "User"
	RestaurantResource          = "Restaurante"
)

type SuccesMessage struct {
	Message Resource
}

func CreateSuccesCreatedMessage(resource Resource) SuccesMessage {
	return SuccesMessage{Message: resource + " created successfully"}
}

func MapErrors(errorMessage string) (int, gin.H) {
	if strings.Contains(errorMessage, "duplicate") {
		return http.StatusConflict, gin.H{"error": errorMessage}
	}

	return http.StatusInternalServerError, gin.H{"error": errorMessage}
}
