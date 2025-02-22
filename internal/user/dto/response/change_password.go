package userresponsedto

type ChangePasswordResDTO struct {
	Message string `json:"message"`
}

func NewChangePasswordResDTO(message string) ChangePasswordResDTO {
	return ChangePasswordResDTO{
		Message: message,
	}
}
