package postgres

import (
	"davidPardoC/rest/common"
	"davidPardoC/rest/domain"
	"errors"
	"net/http"

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

func (repository *UserPostgresRepository) GetUserByEmail(email string) (*domain.User, *common.CustomError) {
	var user domain.User
	result := repository.dB.Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &common.CustomError{StatusCode: http.StatusNotFound, Message: "Email not foud"}
	}
	return &user, nil
}

func (repository *UserPostgresRepository) GetAllUsers() (*[]domain.User, *common.CustomError) {
	var users []domain.User
	result := repository.dB.Find(&users)
	if result.Error != nil {
		return nil, &common.CustomError{StatusCode: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	return &users, nil
}
