package config

import (
	authInfraestructure "davidPardoC/rest/auth/infraestructure"
	"davidPardoC/rest/domain"
	"davidPardoC/rest/users/infraestructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// added cors
	r.Use(cors.Default())

	v1 := r.Group("/api/v1")
	{
		infraestructure.CreateUserRoutes(v1, db)
		authInfraestructure.CreateAuthRoutes(v1, db)
	}
	return r
}

func MakeMigrations(db *gorm.DB) {
	db.AutoMigrate(domain.User{})
}
