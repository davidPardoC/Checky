package infraestructure

import (
	"davidPardoC/rest/roles/adapters"
	"davidPardoC/rest/roles/repository/postgres"
	"davidPardoC/rest/roles/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRolesRouter(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	repository := postgres.NewUserPostgresRepository(db)
	useCases := usecases.NewRolesUseCases(repository)
	rolesAdapters := adapters.NewRolesAdapters(useCases)

	rolesRouter := r.Group("roles")
	{
		rolesRouter.GET("/", func(ctx *gin.Context) {

			roles := rolesAdapters.GetAllRoles()
			ctx.JSON(http.StatusOK, roles)

		})

		rolesRouter.POST("/", func(ctx *gin.Context) {})
	}
	return rolesRouter
}
