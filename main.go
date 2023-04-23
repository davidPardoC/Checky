package main

import (
	"davidPardoC/rest/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.SetupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
