package userrequestdto

type UnBlockUserReqDTO struct {
	UserID *uint `json:"userId"`
}

func NewUnBlockUserReqDTO() *UnBlockUserReqDTO {
	return new(UnBlockUserReqDTO)
}
