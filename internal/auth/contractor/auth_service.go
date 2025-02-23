package authcontractor

import (
	"context"

	authrequestdto "github.com/ladmakhi81/learning-management-system/internal/auth/dto/request"
	userrequestdto "github.com/ladmakhi81/learning-management-system/internal/user/dto/request"
)

type AuthService interface {
	Login(ctx context.Context, dto authrequestdto.LoginReqDTO) (string, error)
	Signup(ctx context.Context, dto userrequestdto.CreateUserReqDTO) (string, error)
}
