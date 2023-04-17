package infraestructure

import (
	"davidPardoC/rest/auth/adapters"
	"davidPardoC/rest/auth/usecases"

	"github.com/gin-gonic/gin"
)

func CreateAuthRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	authUseCases := usecases.NewAuthUseCases()
	authAdapters := adapters.NewAuthAdapters(authUseCases)

	authRouter := r.Group("auth")
	{
		authRouter.POST("/login")
		authRouter.POST("/signup", func(ctx *gin.Context) {
			authAdapters.HandleSingup("testemail.go")
		})
	}

	return authRouter

}
