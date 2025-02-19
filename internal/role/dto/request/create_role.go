package rolerequestdto

import roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"

type CreateRoleReqDTO struct {
	Name        string
	Permissions roleentity.Permissions
}

func NewCreateRoleReqDTO() *CreateRoleReqDTO {
	return new(CreateRoleReqDTO)
}
