package config

import (
	authInfra "davidPardoC/rest/auth/infraestructure"
	"davidPardoC/rest/domain"
	rolesInfra "davidPardoC/rest/roles/infraestructure"
	userInfra "davidPardoC/rest/users/infraestructure"

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
		userInfra.CreateUserRoutes(v1, db)
		authInfra.CreateAuthRoutes(v1, db)
		rolesInfra.CreateRolesRouter(v1, db)
	}
	return r
}

func MakeMigrations(db *gorm.DB) {
	db.AutoMigrate(domain.User{})
	db.AutoMigrate(domain.Role{})
}
