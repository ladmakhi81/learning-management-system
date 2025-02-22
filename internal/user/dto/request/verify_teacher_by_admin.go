package userrequestdto

type VerifyTeacherByAdminReqDTO struct {
	TeacherId uint `json:"teacherId"`
}

func NewVerifyTeacherByAdminReqDTO() *VerifyTeacherByAdminReqDTO {
	return new(VerifyTeacherByAdminReqDTO)
}
