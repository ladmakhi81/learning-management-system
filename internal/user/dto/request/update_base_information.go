package userrequestdto

type UpdateBaseInformationReqBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
}

func NewUpdateBaseInformationReqBody() *UpdateBaseInformationReqBody {
	return new(UpdateBaseInformationReqBody)
}
