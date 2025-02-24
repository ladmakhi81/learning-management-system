package securitymiddleware

import (
	securitycontractor "github.com/ladmakhi81/learning-management-system/internal/security/contractor"
)

type Middleware struct {
	tokenSvc securitycontractor.TokenService
}

func NewMiddleware(
	tokenSvc securitycontractor.TokenService,
) Middleware {
	return Middleware{
		tokenSvc: tokenSvc,
	}
}
