package userrequestdto

type AssignRoleReqDTO struct {
	RoleID uint `json:"roleId"`
	UserID uint `json:"userId"`
}

func NewAssignRoleReqDTO() *AssignRoleReqDTO {
	return new(AssignRoleReqDTO)
}
