package authcontractor

import (
	"context"

	authrequestdto "github.com/ladmakhi81/learning-management-system/internal/auth/dto/request"
)

type SessionService interface {
	StoreSession(ctx context.Context, dto authrequestdto.SessionDTO) error
	GetSessionByUserId(ctx context.Context, userId uint) (*authrequestdto.SessionDTO, error)
}
