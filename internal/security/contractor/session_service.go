package securitycontractor

import (
	"context"

	securitytype "github.com/ladmakhi81/learning-management-system/internal/security/type"
)

type SessionService interface {
	StoreSession(ctx context.Context, dto securitytype.SessionDTO) error
	GetSessionByUserId(ctx context.Context, userId uint) (*securitytype.SessionDTO, error)
}
