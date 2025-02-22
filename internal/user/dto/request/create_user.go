package userrequestdto

type CreateUserReqDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

func NewCreateUserReqDTO() *CreateUserReqDTO {
	return new(CreateUserReqDTO)
}
