package main

import (
	"davidPardoC/rest/common"
	"davidPardoC/rest/config"
)

func main() {
	database, _ := common.GetDatabase()
	r := config.SetupRouter(database)
	config.MakeMigrations(database)
	r.Run() // listen and serve on 0.0.0.0:8080
}
