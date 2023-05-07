package repository

import "davidPardoC/rest/domain"

type RoleRepository interface {
	GetAllRoles() []domain.Role
	InsertRole(name string) (int64, error)
}
