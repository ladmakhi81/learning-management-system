package authcontractor

import authrequestdto "github.com/ladmakhi81/learning-management-system/internal/auth/dto/request"

type TokenService interface {
	GenerateToken(claim authrequestdto.GenerateTokenDTO) (string, error)
}
