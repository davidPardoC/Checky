package postgres

import (
	"davidPardoC/rest/domain"

	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	dB *gorm.DB
}

func NewUserPostgresRepository(dB *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{dB: dB}
}

func (repository *UserPostgresRepository) InsertNewUser(user domain.User) (int64, error) {
	result := repository.dB.Create(&user)
	return result.RowsAffected, result.Error
}

func (upr *UserPostgresRepository) GetUserByEmail(email string) domain.User {
	return domain.User{}
}

func (upr *UserPostgresRepository) GetAllUsers() []domain.User {
	return []domain.User{}
}
