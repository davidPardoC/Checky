package infraestructure

import (
	"davidPardoC/rest/common"
	"davidPardoC/rest/users/adapters"
	"davidPardoC/rest/users/repository/postgres"
	"davidPardoC/rest/users/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUserRoutes(r *gin.RouterGroup, db *gorm.DB) {

	postgresUserRepository := postgres.NewUserPostgresRepository(db)
	useruseCase := usecase.NewUserUseCase(postgresUserRepository)
	userAdapter := adapters.NewUserAdapter(useruseCase)

	userRouter := r.Group("/users")
	{
		userRouter.GET("/", func(ctx *gin.Context) {
			userAdapter.HandleNewUser()
			message := common.CreateSuccesCreatedMessage(common.UserResource)
			ctx.JSON(200, message)
		})
		userRouter.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			userAdapter.HandleGetUserById(id)
		})

		userRouter.POST("/", func(ctx *gin.Context) {})

		userRouter.PUT("/", func(ctx *gin.Context) {})

		userRouter.DELETE("/", func(ctx *gin.Context) {})

	}
}
