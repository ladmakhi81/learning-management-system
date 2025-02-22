package userresponsedto

type UnBlockUserResDTO struct {
	Message string `json:"message"`
}

func NewUnBlockUserResDTO(message string) UnBlockUserResDTO {
	return UnBlockUserResDTO{
		Message: message,
	}
}
