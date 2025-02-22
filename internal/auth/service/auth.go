package authservice

import (
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
	userSvc  usercontractor.UserService
	tokenSvc authcontractor.TokenService
}

func NewAuthServiceImpl(
	userSvc usercontractor.UserService,
	tokenSvc authcontractor.TokenService,
) AuthServiceImpl {
	return AuthServiceImpl{
		userSvc:  userSvc,
		tokenSvc: tokenSvc,
	}
}

func (authSvc AuthServiceImpl) Login(dto authrequestdto.LoginReqDTO) (string, error) {
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
	// save on the redis
	return accessToken, nil
}

func (authSvc AuthServiceImpl) Signup(dto userrequestdto.CreateUserReqDTO) (string, error) {
	user, userErr := authSvc.userSvc.CreateUser(dto)
	if userErr != nil {
		return "", userErr
	}
	claim := authrequestdto.NewGenerateTokenDTO(user.ID, user.RoleID)
	accessToken, accessTokenErr := authSvc.tokenSvc.GenerateToken(claim)
	if accessTokenErr != nil {
		return "", accessTokenErr
	}
	// save on the redis
	return accessToken, nil
}
