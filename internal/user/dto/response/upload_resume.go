package userresponsedto

type UploadResumeResDTO struct {
	ResumeFileURL string `json:"resumeFileURL"`
}

func NewUploadResumeResDTO(resumeFileURL string) UploadResumeResDTO {
	return UploadResumeResDTO{
		ResumeFileURL: resumeFileURL,
	}
}
