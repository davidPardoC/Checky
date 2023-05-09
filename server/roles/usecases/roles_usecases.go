package usecases

import (
	"davidPardoC/rest/domain"
	"davidPardoC/rest/roles/repository"
	"fmt"
)

type RolesUseCases struct {
	repository repository.RoleRepository
}

func NewRolesUseCases(repository repository.RoleRepository) *RolesUseCases {
	useCases := &RolesUseCases{repository: repository}
	useCases.createInitialRoles()
	return useCases
}

func (useCase *RolesUseCases) GetAllRoles() []domain.Role {
	return useCase.repository.GetAllRoles()
}

func (useCase *RolesUseCases) CreateRole(role domain.Role) (int64, error) {
	return useCase.repository.InsertRole(role.Name)

}

func (useCase *RolesUseCases) createInitialRoles() {
	adminRole := domain.Role{Name: "admin"}
	cook := domain.Role{Name: "cook"}
	waiter := domain.Role{Name: "waiter"}
	assistant := domain.Role{Name: "assistant"}
	roles := []domain.Role{adminRole, cook, waiter, assistant}
	for i, role := range roles {
		fmt.Printf("[%d] Created role: %s\n", i, role.Name)
		useCase.CreateRole(role)
	}
}
