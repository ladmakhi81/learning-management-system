package authhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authconstant "github.com/ladmakhi81/learning-management-system/internal/auth/constant"
	authcontractor "github.com/ladmakhi81/learning-management-system/internal/auth/contractor"
	authrequestdto "github.com/ladmakhi81/learning-management-system/internal/auth/dto/request"
	authresponsedto "github.com/ladmakhi81/learning-management-system/internal/auth/dto/response"
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	basehandler "github.com/ladmakhi81/learning-management-system/internal/base/handler"
	userrequestdto "github.com/ladmakhi81/learning-management-system/internal/user/dto/request"
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
	dto := authrequestdto.NewLoginReqDTO()
	if err := ctx.Bind(dto); err != nil {
		return nil, baseerror.NewClientErr(
			authconstant.INVALID_REQUEST_BODY,
			http.StatusBadRequest,
		)
	}
	generatedAccessToken, authErr := h.authSvc.Login(ctx.Request.Context(), *dto)
	if authErr != nil {
		return nil, authErr
	}
	res := authresponsedto.NewLoginResDTO(generatedAccessToken)
	return basehandler.NewResponse(res, http.StatusOK), nil
}

func (h AuthHandler) Signup(ctx *gin.Context) (*basehandler.Response, error) {
	dto := userrequestdto.NewCreateUserReqDTO()
	if err := ctx.Bind(dto); err != nil {
		return nil, baseerror.NewClientErr(
			authconstant.INVALID_REQUEST_BODY,
			http.StatusBadRequest,
		)
	}
	generatedAccessToken, authErr := h.authSvc.Signup(ctx.Request.Context(), *dto)
	if authErr != nil {
		return nil, authErr
	}
	res := authresponsedto.NewSignupResDTO(generatedAccessToken)
	return basehandler.NewResponse(res, http.StatusCreated), nil
}
