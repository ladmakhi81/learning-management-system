package securitymiddleware

import (
	securitycontractor "github.com/ladmakhi81/learning-management-system/internal/security/contractor"
)

type Middleware struct {
	tokenSvc   securitycontractor.TokenService
	sessionSvc securitycontractor.SessionService
}

func NewMiddleware(
	tokenSvc securitycontractor.TokenService,
	sessionSvc securitycontractor.SessionService,
) Middleware {
	return Middleware{
		tokenSvc:   tokenSvc,
		sessionSvc: sessionSvc,
	}
}
