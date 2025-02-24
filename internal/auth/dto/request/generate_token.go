package authrequestdto

type TokenDTO struct {
	UserID uint
	RoleID *uint
}

func NewTokenDTO(userId uint, roleId *uint) TokenDTO {
	return TokenDTO{
		UserID: userId,
		RoleID: roleId,
	}
}
