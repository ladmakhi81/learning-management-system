package securityservice

import (
	"context"
	"encoding/json"

	authconstant "github.com/ladmakhi81/learning-management-system/internal/auth/constant"
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	securitytype "github.com/ladmakhi81/learning-management-system/internal/security/type"
	pkgredisclient "github.com/ladmakhi81/learning-management-system/pkg/redis-client"
)

type SessionServiceImpl struct {
	redisClient *pkgredisclient.RedisClient
}

func NewSessionServiceImpl(
	redisClient *pkgredisclient.RedisClient,
) SessionServiceImpl {
	return SessionServiceImpl{
		redisClient: redisClient,
	}
}

func (sessionSvc SessionServiceImpl) StoreSession(ctx context.Context, dto securitytype.SessionDTO) error {
	err := sessionSvc.redisClient.SetHashValue(ctx, authconstant.AUTH_SESSION_REDIS_KEY, string(dto.UserId), dto)
	if err != nil {
		return baseerror.NewServerErr(
			err.Error(),
			"SessionServiceImpl.StoreSession",
		)
	}
	return nil
}

func (sessionSvc SessionServiceImpl) GetSessionByUserId(ctx context.Context, userId uint) (*securitytype.SessionDTO, error) {
	data, dataErr := sessionSvc.redisClient.GetHashValue(ctx, authconstant.AUTH_SESSION_REDIS_KEY, string(userId))
	if dataErr != nil {
		return nil, baseerror.NewServerErr(
			dataErr.Error(),
			"SessionServiceImpl.GetSessionByUserId",
		)
	}
	session := new(securitytype.SessionDTO)
	if err := json.Unmarshal([]byte(data), session); err != nil {
		return nil, baseerror.NewServerErr(
			err.Error(),
			"SessionServiceImpl.GetSessionByUserId",
		)
	}
	return session, nil
}
