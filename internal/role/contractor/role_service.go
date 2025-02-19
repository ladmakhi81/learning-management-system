package rolecontractor

import (
	basetype "github.com/ladmakhi81/learning-management-system/internal/base/type"
	rolerequestdto "github.com/ladmakhi81/learning-management-system/internal/role/dto/request"
	roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"
)

type RoleService interface {
	CreateRole(dto *rolerequestdto.CreateRoleReqDTO) (*roleentity.Role, error)
	DeleteRoleById(id uint) error
	FindRoleById(id uint) (*roleentity.Role, error)
	FindRoleByName(name string) (*roleentity.Role, error)
	GetRoles(page, limit int) ([]roleentity.Role, error)
	GetRolesPaginationMetadata(currentPage uint, limit uint) (*basetype.PaginationMetadata, error)
}
