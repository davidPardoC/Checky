package auth_test

import (
	"davidPardoC/rest/common"
	"davidPardoC/rest/config"
	"davidPardoC/rest/domain"
)

var basePath string = "/api/v1"

var database, _ = common.GetDatabase()

var router = config.SetupRouter(database)

func SetupEmptyUserDatabaseTests() {
	database.Migrator().DropTable(domain.User{})
	config.MakeMigrations(database)
}

func GetFinalEndpoint(path string) string {
	return basePath + "/auth" + path
}
