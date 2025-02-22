package rolecontractor

import (
	roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"
)

type RoleRepository interface {
	CreateRole(role *roleentity.Role) error
	DeleteRoleById(id uint) error
	FindRoleById(id uint) (*roleentity.Role, error)
	FindRoleByName(name string) (*roleentity.Role, error)
	GetRoles(page, limit int) ([]roleentity.Role, error)
	GetRolesCount() (int, error)
}
