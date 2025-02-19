package rolemapper

import (
	"fmt"

	roleresponsedto "github.com/ladmakhi81/learning-management-system/internal/role/dto/response"
	roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"
)

type RoleMapper struct{}

func NewRoleMapper() RoleMapper {
	return RoleMapper{}
}

func (m RoleMapper) MapRoleToRoleResponseDTO(role *roleentity.Role) roleresponsedto.RoleResDTO {
	dto := roleresponsedto.NewRoleResDTO()
	dto.Name = role.Name
	dto.ID = role.ID
	dto.CreatedAt = &role.CreatedAt
	dto.CreatedByID = role.CreatedByID
	dto.Lock = role.Lock
	dto.Permissions = role.Permissions
	dto.UpdatedAt = &role.UpdatedAt
	return dto
}

func (m RoleMapper) MapRolesToRolesResponseDTO(roles []roleentity.Role) []roleresponsedto.RoleResDTO {
	fmt.Println(len(roles))
	dtos := make([]roleresponsedto.RoleResDTO, len(roles))
	for index, role := range roles {
		dtos[index] = m.MapRoleToRoleResponseDTO(&role)
	}
	return dtos
}
