package common

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=checky port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}

func GetTestDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=checky_test port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}
