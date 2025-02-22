package authcontractor

import (
	authrequestdto "github.com/ladmakhi81/learning-management-system/internal/auth/dto/request"
	userrequestdto "github.com/ladmakhi81/learning-management-system/internal/user/dto/request"
)

type AuthService interface {
	Login(dto authrequestdto.LoginReqDTO) (string, error)
	Signup(dto userrequestdto.CreateUserReqDTO) (string, error)
}
