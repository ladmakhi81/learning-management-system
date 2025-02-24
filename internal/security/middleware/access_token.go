package securitymiddleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	basetype "github.com/ladmakhi81/learning-management-system/internal/base/type"
	securityconstant "github.com/ladmakhi81/learning-management-system/internal/security/constant"
)

func (m Middleware) CheckAccessToken(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")
	authorizationSegments := strings.Split(authorization, " ")
	if len(authorizationSegments) != 2 {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			basetype.NewUnauthorizedResponse(),
		)
		return
	}
	bearer := authorizationSegments[0]
	if strings.ToLower(bearer) != "bearer" {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			basetype.NewUnauthorizedResponse(),
		)
		return
	}
	token := authorizationSegments[1]
	if token == "" {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			basetype.NewUnauthorizedResponse(),
		)
		return
	}
	verifiedToken, verifiedTokenErr := m.tokenSvc.VerifyToken(token)
	if verifiedTokenErr != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			basetype.NewUnauthorizedResponse(),
		)
		return
	}
	session, sessionErr := m.sessionSvc.GetSessionByUserId(ctx.Request.Context(), verifiedToken.UserID)
	if sessionErr != nil {
		fmt.Println("Session Error: ", sessionErr)
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			basetype.NewUnauthorizedResponse(),
		)
		return
	}
	if session == nil || session.AccessToken != token {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			basetype.NewUnauthorizedResponse(),
		)
		return
	}
	ctx.Set(securityconstant.SECURITY_DECODE_KEY, session)
	ctx.Next()
}
