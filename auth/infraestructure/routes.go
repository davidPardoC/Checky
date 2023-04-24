package infraestructure

import (
	"davidPardoC/rest/auth/adapters"
	"davidPardoC/rest/auth/dtos"
	"davidPardoC/rest/auth/usecases"
	"davidPardoC/rest/users/repository/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAuthRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	userRepository := postgres.NewUserPostgresRepository(db)
	authUseCases := usecases.NewAuthUseCases(userRepository)
	authAdapters := adapters.NewAuthAdapters(authUseCases)

	authRouter := r.Group("auth")
	{
		authRouter.POST("/login", func(ctx *gin.Context) {
			var json dtos.LoginDto

			if err := ctx.ShouldBindJSON(&json); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			res := authAdapters.HandleLogin(json)
			ctx.String(http.StatusOK, res)
		})
		authRouter.POST("/signup", func(ctx *gin.Context) {
			authAdapters.HandleSingup("testemail.go")
		})
	}

	return authRouter

}
