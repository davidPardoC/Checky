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
	employeeRole := domain.Role{Name: "employee"}
	roles := []domain.Role{adminRole, employeeRole}
	for i, role := range roles {
		fmt.Printf("[%d] Created role: %s\n", i, role.Name)
		useCase.CreateRole(role)
	}
}
