package main

import (
	"davidPardoC/rest/common"
	"davidPardoC/rest/config"

	"github.com/gin-gonic/gin"
)

func main() {
	database, _ := common.GetDatabase()
	r := gin.Default()
	config.SetupRoutes(r, database)
	config.MakeMigrations(database)
	r.Run() // listen and serve on 0.0.0.0:8080
}
