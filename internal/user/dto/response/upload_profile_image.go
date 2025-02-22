package userresponsedto

type UploadProfileImageResDTO struct {
	ProfileFileName string `json:"profileFileName"`
}

func NewUploadProfileImageResDTO(profileFileName string) UploadProfileImageResDTO {
	return UploadProfileImageResDTO{ProfileFileName: profileFileName}
}
