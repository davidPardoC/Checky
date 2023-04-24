package common

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}
