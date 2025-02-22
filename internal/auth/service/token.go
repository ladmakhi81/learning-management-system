package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	authrequestdto "github.com/ladmakhi81/learning-management-system/internal/auth/dto/request"
	baseconfig "github.com/ladmakhi81/learning-management-system/internal/base/config"
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	basetype "github.com/ladmakhi81/learning-management-system/internal/base/type"
)

type TokenServiceImpl struct {
	config *baseconfig.Config
}

func NewTokenServiceImpl(
	config *baseconfig.Config,
) TokenServiceImpl {
	return TokenServiceImpl{
		config: config,
	}
}

func (tokenSvc TokenServiceImpl) GenerateToken(claim authrequestdto.GenerateTokenDTO) (string, error) {
	jwtClaim := basetype.NewAuthClaim(claim.UserID, claim.RoleID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	secretKey := tokenSvc.config.SecretKey
	signedToken, signedTokenErr := token.SignedString([]byte(secretKey))
	if signedTokenErr != nil {
		return "", baseerror.NewServerErr(
			signedTokenErr.Error(),
			"TokenServiceImpl.GenerateToken",
		)
	}
	return signedToken, nil
}
