package userresponsedto

type CreateUserResDTO struct {
	User UserResDTO `json:"user"`
}

func NewCreateUserResDTO(user UserResDTO) CreateUserResDTO {
	return CreateUserResDTO{
		User: user,
	}
}
