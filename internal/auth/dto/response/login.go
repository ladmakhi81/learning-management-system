package authresponsedto

type LoginResDTO struct {
	AccessToken string `json:"accessToken"`
}

func NewLoginResDTO(accessToken string) LoginResDTO {
	return LoginResDTO{
		AccessToken: accessToken,
	}
}
