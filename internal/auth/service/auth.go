package authservice

import (
	"context"
	"net/http"

	authconstant "github.com/ladmakhi81/learning-management-system/internal/auth/constant"
	authcontractor "github.com/ladmakhi81/learning-management-system/internal/auth/contractor"
	authrequestdto "github.com/ladmakhi81/learning-management-system/internal/auth/dto/request"
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	usercontractor "github.com/ladmakhi81/learning-management-system/internal/user/contractor"
	userrequestdto "github.com/ladmakhi81/learning-management-system/internal/user/dto/request"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userSvc    usercontractor.UserService
	tokenSvc   authcontractor.TokenService
	sessionSvc authcontractor.SessionService
}

func NewAuthServiceImpl(
	userSvc usercontractor.UserService,
	tokenSvc authcontractor.TokenService,
	sessionSvc authcontractor.SessionService,
) AuthServiceImpl {
	return AuthServiceImpl{
		userSvc:    userSvc,
		tokenSvc:   tokenSvc,
		sessionSvc: sessionSvc,
	}
}

func (authSvc AuthServiceImpl) Login(ctx context.Context, dto authrequestdto.LoginReqDTO) (string, error) {
	user, userErr := authSvc.userSvc.FindUserByPhone(dto.Phone)
	if userErr != nil {
		return "", userErr
	}
	if user == nil {
		return "", baseerror.NewClientErr(
			authconstant.INVALID_PHONE_PASSWORD,
			http.StatusNotFound,
		)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return "", baseerror.NewClientErr(
			authconstant.INVALID_PHONE_PASSWORD,
			http.StatusNotFound,
		)
	}
	claim := authrequestdto.NewGenerateTokenDTO(user.ID, user.RoleID)
	accessToken, accessTokenErr := authSvc.tokenSvc.GenerateToken(claim)
	if accessTokenErr != nil {
		return "", accessTokenErr
	}
	session := authrequestdto.NewSessionDTO(user.ID, accessToken)
	if err := authSvc.sessionSvc.StoreSession(ctx, session); err != nil {
		return "", err
	}
	return accessToken, nil
}

func (authSvc AuthServiceImpl) Signup(ctx context.Context, dto userrequestdto.CreateUserReqDTO) (string, error) {
	user, userErr := authSvc.userSvc.CreateUser(dto)
	if userErr != nil {
		return "", userErr
	}
	claim := authrequestdto.NewGenerateTokenDTO(user.ID, user.RoleID)
	accessToken, accessTokenErr := authSvc.tokenSvc.GenerateToken(claim)
	if accessTokenErr != nil {
		return "", accessTokenErr
	}
	session := authrequestdto.NewSessionDTO(user.ID, accessToken)
	if err := authSvc.sessionSvc.StoreSession(ctx, session); err != nil {
		return "", err
	}
	return accessToken, nil
}
