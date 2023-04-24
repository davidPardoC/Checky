package config

import (
	authInfraestructure "davidPardoC/rest/auth/infraestructure"
	"davidPardoC/rest/users/infraestructure"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	v1 := r.Group("/api/v1")
	{
		infraestructure.CreateUserRoutes(v1, db)
		authInfraestructure.CreateAuthRoutes(v1, db)
	}
}
