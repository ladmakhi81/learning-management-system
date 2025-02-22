package userresponsedto

type BlockUserResDTO struct {
	Message string `json:"message"`
}

func NewBlockUserResDTO(message string) BlockUserResDTO {
	return BlockUserResDTO{Message: message}
}
