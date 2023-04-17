package config

import (
	authInfraestructure "davidPardoC/rest/auth/infraestructure"
	"davidPardoC/rest/users/infraestructure"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		infraestructure.CreateUserRoutes(v1)
		authInfraestructure.CreateAuthRoutes(v1)
	}
}
