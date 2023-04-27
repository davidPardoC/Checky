package infraestructure

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUserRoutes(r *gin.RouterGroup, db *gorm.DB) {

	userRouter := r.Group("/users")
	{
		userRouter.GET("/", func(ctx *gin.Context) {

		})
		userRouter.GET("/:id", func(ctx *gin.Context) {

		})

		userRouter.POST("/", func(ctx *gin.Context) {})

		userRouter.PUT("/", func(ctx *gin.Context) {})

		userRouter.DELETE("/", func(ctx *gin.Context) {})

	}
}
