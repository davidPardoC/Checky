package main

import (
	"davidPardoC/rest/common"
	"davidPardoC/rest/config"
	"fmt"

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
	fmt.Println("Hola")
	r.Run() // listen and serve on 0.0.0.0:8080
}
