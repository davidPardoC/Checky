package postgres

import (
	"davidPardoC/rest/domain"

	"gorm.io/gorm"
)

type RolesPostgresRepository struct {
	dB *gorm.DB
}

func NewUserPostgresRepository(dB *gorm.DB) *RolesPostgresRepository {
	return &RolesPostgresRepository{dB: dB}
}

func (repo *RolesPostgresRepository) GetAllRoles() []domain.Role {
	var roles = []domain.Role{}
	repo.dB.Find(&roles)
	return roles
}

func (repo *RolesPostgresRepository) InsertRole(name string) (int64, error) {
	role := domain.Role{Name: name}
	result := repo.dB.Create(&role)
	return result.RowsAffected, result.Error
}
