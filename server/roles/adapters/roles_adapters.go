package adapters

import (
	"davidPardoC/rest/domain"
	"davidPardoC/rest/roles/usecases"
)

type RolesAdapters struct {
	useCases *usecases.RolesUseCases
}

func NewRolesAdapters(useCases *usecases.RolesUseCases) *RolesAdapters {
	return &RolesAdapters{useCases}
}

func (adapter *RolesAdapters) GetAllRoles() []domain.Role {
	return adapter.useCases.GetAllRoles()
}
