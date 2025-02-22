package userrequestdto

type ChangePasswordReqDTO struct {
	UserID   uint   `json:"userId"`
	Password string `json:"password"`
}

func NewChangePasswordReqDTO() *ChangePasswordReqDTO {
	return new(ChangePasswordReqDTO)
}
