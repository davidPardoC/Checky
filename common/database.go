package common

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase() (*gorm.DB, error) {
	env := os.Getenv("GO_ENV")
	var dsn string = "host=localhost user=postgres password=mysecretpassword dbname=checky port=5432 sslmode=disable"
	if env == "test" {
		dsn = "host=localhost user=postgres password=mysecretpassword dbname=checky_test port=5433 sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
