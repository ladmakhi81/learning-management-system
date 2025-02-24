package securitycontractor

import (
	authrequestdto "github.com/ladmakhi81/learning-management-system/internal/auth/dto/request"
	basetype "github.com/ladmakhi81/learning-management-system/internal/base/type"
)

type TokenService interface {
	GenerateToken(claim authrequestdto.TokenDTO) (string, error)
	VerifyToken(token string) (*basetype.AuthClaim, error)
}
