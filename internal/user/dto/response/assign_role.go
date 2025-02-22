package userresponsedto

type AssignRoleResDTO struct {
	Message string `json:"message"`
}

func NewAssignRoleResDTO(message string) AssignRoleResDTO {
	return AssignRoleResDTO{
		Message: message,
	}
}
