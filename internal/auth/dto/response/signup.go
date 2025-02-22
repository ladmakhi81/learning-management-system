package authresponsedto

type SignupResDTO struct {
	AccessToken string `json:"accessToken"`
}

func NewSignupResDTO(accessToken string) SignupResDTO {
	return SignupResDTO{
		AccessToken: accessToken,
	}
}
