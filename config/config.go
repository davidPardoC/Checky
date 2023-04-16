package config

import (
	"davidPardoC/rest/infraestructure"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		infraestructure.CreateUserRoutes(v1)
	}
}
