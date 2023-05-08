package infraestructure

import (
	"davidPardoC/rest/users/adapters"
	"davidPardoC/rest/users/repository/postgres"
	"davidPardoC/rest/users/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUserRoutes(r *gin.RouterGroup, db *gorm.DB) {

	repository := postgres.NewUserPostgresRepository(db)
	useCases := usecases.NewUserUseCase(repository)
	adapters := adapters.NewUserAdapter(useCases)

	userRouter := r.Group("/users")
	{
		userRouter.GET("/", func(ctx *gin.Context) {
			users, err := adapters.GetAllUsers()

			if err != nil {
				ctx.JSON(err.StatusCode, err.Message)
				return
			}

			ctx.JSON(http.StatusOK, users)
		})
		userRouter.GET("/:id", func(ctx *gin.Context) {

		})

		userRouter.POST("/", func(ctx *gin.Context) {})

		userRouter.PUT("/", func(ctx *gin.Context) {})

		userRouter.DELETE("/", func(ctx *gin.Context) {})

	}
}
