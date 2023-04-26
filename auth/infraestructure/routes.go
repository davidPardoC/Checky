package infraestructure

import (
	"davidPardoC/rest/auth/adapters"
	"davidPardoC/rest/auth/dtos"
	"davidPardoC/rest/auth/usecases"
	"davidPardoC/rest/common"
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
				return
			}

			res := authAdapters.HandleLogin(json)
			ctx.JSON(http.StatusOK, gin.H{"message": res})
		})

		authRouter.POST("/signup", func(ctx *gin.Context) {
			var json dtos.SignupDto

			if err := ctx.ShouldBindJSON(&json); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			message, err := authAdapters.HandleSingup(json)
			if err != nil {
				ctx.JSON(common.MapErrors(err.Error()))
				return
			}
			ctx.JSON(http.StatusOK, message)
		})
	}

	return authRouter

}
