package authrequestdto

type GenerateTokenDTO struct {
	UserID uint
	RoleID *uint
}

func NewGenerateTokenDTO(userId uint, roleId *uint) GenerateTokenDTO {
	return GenerateTokenDTO{
		UserID: userId,
		RoleID: roleId,
	}
}
