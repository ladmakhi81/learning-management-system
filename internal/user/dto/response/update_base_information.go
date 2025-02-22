package userresponsedto

type UpdateBaseInformationResDTO struct {
	Message string `json:"message"`
}

func NewUpdateBaseInformationResDTO(message string) UpdateBaseInformationResDTO {
	return UpdateBaseInformationResDTO{
		Message: message,
	}
}
