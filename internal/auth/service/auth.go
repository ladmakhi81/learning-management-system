package authservice

import (
	"context"
	"net/http"

	authconstant "github.com/ladmakhi81/learning-management-system/internal/auth/constant"
	authrequestdto "github.com/ladmakhi81/learning-management-system/internal/auth/dto/request"
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	rolecontractor "github.com/ladmakhi81/learning-management-system/internal/role/contractor"
	roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"
	securitycontractor "github.com/ladmakhi81/learning-management-system/internal/security/contractor"
	securitytype "github.com/ladmakhi81/learning-management-system/internal/security/type"
	usercontractor "github.com/ladmakhi81/learning-management-system/internal/user/contractor"
	userrequestdto "github.com/ladmakhi81/learning-management-system/internal/user/dto/request"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userSvc    usercontractor.UserService
	tokenSvc   securitycontractor.TokenService
	sessionSvc securitycontractor.SessionService
	roleSvc    rolecontractor.RoleService
}

func NewAuthServiceImpl(
	userSvc usercontractor.UserService,
	tokenSvc securitycontractor.TokenService,
	sessionSvc securitycontractor.SessionService,
	roleSvc rolecontractor.RoleService,
) AuthServiceImpl {
	return AuthServiceImpl{
		userSvc:    userSvc,
		tokenSvc:   tokenSvc,
		sessionSvc: sessionSvc,
		roleSvc:    roleSvc,
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
	var userRoleId *uint
	var userPermissions roleentity.Permissions
	if user.RoleID != nil {
		role, roleErr := authSvc.roleSvc.FindRoleById(*user.RoleID)
		if roleErr != nil {
			return "", roleErr
		}
		userRoleId = &role.ID
		userPermissions = role.Permissions
	}

	claim := authrequestdto.NewTokenDTO(user.ID, user.RoleID)
	accessToken, accessTokenErr := authSvc.tokenSvc.GenerateToken(claim)
	if accessTokenErr != nil {
		return "", accessTokenErr
	}

	session := securitytype.NewSessionDTO(user.ID, accessToken, userRoleId, userPermissions)
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
	claim := authrequestdto.NewTokenDTO(user.ID, user.RoleID)
	accessToken, accessTokenErr := authSvc.tokenSvc.GenerateToken(claim)
	if accessTokenErr != nil {
		return "", accessTokenErr
	}
	session := securitytype.NewSessionDTO(user.ID, accessToken, nil, nil)
	if err := authSvc.sessionSvc.StoreSession(ctx, session); err != nil {
		return "", err
	}
	return accessToken, nil
}
