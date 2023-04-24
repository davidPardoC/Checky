package infraestructure

import (
	"davidPardoC/rest/auth/adapters"
	"davidPardoC/rest/auth/dtos"
	"davidPardoC/rest/auth/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAuthRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	authUseCases := usecases.NewAuthUseCases()
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
