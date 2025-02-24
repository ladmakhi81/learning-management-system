package securitytype

import roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"

type SessionDTO struct {
	UserId      uint
	AccessToken string
	RoleID      *uint
	Permissions roleentity.Permissions
}

func NewSessionDTO(
	userId uint,
	accessToken string,
	roleID *uint,
	permissions roleentity.Permissions,
) SessionDTO {
	return SessionDTO{
		UserId:      userId,
		AccessToken: accessToken,
		RoleID:      roleID,
		Permissions: permissions,
	}
}
