package infraestructure

import (
	"davidPardoC/rest/common"

	"github.com/gin-gonic/gin"
)

func CreateUserRoutes(r *gin.RouterGroup) {
	userRouter := r.Group("/users")
	{
		userRouter.GET("/", func(ctx *gin.Context) {
			message := common.CreateSuccesCreatedMessage(common.UserResource)
			ctx.JSON(200, message)
		})
		userRouter.POST("/", func(ctx *gin.Context) {})
		userRouter.PUT("/", func(ctx *gin.Context) {})
		userRouter.DELETE("/", func(ctx *gin.Context) {})

	}
}
