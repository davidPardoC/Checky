package postgres

import (
	"davidPardoC/rest/users/domain"

	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	dB *gorm.DB
}

func NewUserPostgresRepository(dB *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{dB: dB}
}

func (upr *UserPostgresRepository) InsertNewUser() string {
	return "User inserted"
}

func (upr *UserPostgresRepository) GetUserByEmail(email string) domain.User {
	return domain.User{}
}

func (upr *UserPostgresRepository) GetAllUsers() []domain.User {
	return []domain.User{}
}
