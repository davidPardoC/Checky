package main

import (
	"davidPardoC/rest/common"
	"davidPardoC/rest/config"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Not .env file")
	}
	database, _ := common.GetDatabase()
	r := config.SetupRouter(database)
	config.MakeMigrations(database)
	r.Run() // listen and serve on 0.0.0.0:8080
}