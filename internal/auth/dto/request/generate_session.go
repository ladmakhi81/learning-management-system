package authrequestdto

type SessionDTO struct {
	UserId      uint
	AccessToken string
}

func NewSessionDTO(userId uint, accessToken string) SessionDTO {
	return SessionDTO{
		UserId:      userId,
		AccessToken: accessToken,
	}
}
