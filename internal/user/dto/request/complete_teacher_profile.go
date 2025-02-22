package userrequestdto

type CompleteTeacherProfileReqDTO struct {
	Bio          string `json:"bio"`
	ProfileImage string `json:"profileImage"`
	ResumeFile   string `json:"resumeFile"`
	Email        string `json:"email"`
	NationalID   string `json:"nationalId"`
}

func NewCompleteTeacherProfileReqDTO() *CompleteTeacherProfileReqDTO {
	return new(CompleteTeacherProfileReqDTO)
}
