package userresponsedto

type CompleteTeacherProfileResDTO struct {
	Message string `json:"message"`
}

func NewCompleteTeacherProfileResDTO(message string) CompleteTeacherProfileResDTO {
	return CompleteTeacherProfileResDTO{
		Message: message,
	}
}
