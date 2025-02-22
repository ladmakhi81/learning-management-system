package userresponsedto

type VerifyTeacherByAdminResDTO struct {
	Message string `json:"message"`
}

func NewVerifyTeacherByAdminResDTO(message string) VerifyTeacherByAdminResDTO {
	return VerifyTeacherByAdminResDTO{
		Message: message,
	}
}
