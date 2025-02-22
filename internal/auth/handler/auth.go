package authhandler

import (
	"github.com/gin-gonic/gin"
	authcontractor "github.com/ladmakhi81/learning-management-system/internal/auth/contractor"
	basehandler "github.com/ladmakhi81/learning-management-system/internal/base/handler"
)

type AuthHandler struct {
	authSvc authcontractor.AuthService
}

func NewAuthHandler(
	authSvc authcontractor.AuthService,
) AuthHandler {
	return AuthHandler{
		authSvc: authSvc,
	}
}

func (h AuthHandler) Login(ctx *gin.Context) (*basehandler.Response, error) {
	return nil, nil
}

func (h AuthHandler) Signup(ctx *gin.Context) (*basehandler.Response, error) {
	return nil, nil
}
