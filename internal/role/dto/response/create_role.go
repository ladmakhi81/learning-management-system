package roleresponsedto

type CreateRoleRes struct {
	Role RoleResDTO `json:"role"`
}

func NewCreateRoleRes(role RoleResDTO) CreateRoleRes {
	return CreateRoleRes{
		Role: role,
	}
}
