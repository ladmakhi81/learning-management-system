package securityservice

import (
	"errors"
	"fmt"
	"time"

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

func (tokenSvc TokenServiceImpl) GenerateToken(claim authrequestdto.TokenDTO) (string, error) {
	jwtClaim := basetype.NewAuthClaim(claim.UserID, claim.RoleID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	secretKey := tokenSvc.getSecretKey()
	signedToken, signedTokenErr := token.SignedString(secretKey)
	if signedTokenErr != nil {
		return "", baseerror.NewServerErr(
			signedTokenErr.Error(),
			"TokenServiceImpl.GenerateToken",
		)
	}
	return signedToken, nil
}

func (tokenSvc TokenServiceImpl) VerifyToken(token string) (*basetype.AuthClaim, error) {
	decodedToken, decodedTokenErr := jwt.ParseWithClaims(token, &basetype.AuthClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method : %v", t.Header["alg"])
		}
		return tokenSvc.getSecretKey(), nil
	})
	if decodedTokenErr != nil || !decodedToken.Valid {
		return nil, errors.New("invalid token")
	}
	claims := decodedToken.Claims.(*basetype.AuthClaim)
	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("expired token")
	}
	return claims, nil
}

func (tokenSvc TokenServiceImpl) getSecretKey() []byte {
	return []byte(tokenSvc.config.SecretKey)
}
