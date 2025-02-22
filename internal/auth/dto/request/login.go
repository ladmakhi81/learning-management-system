package authrequestdto

type LoginReqDTO struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func NewLoginReqDTO() *LoginReqDTO {
	return new(LoginReqDTO)
}
