package securitymiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	basetype "github.com/ladmakhi81/learning-management-system/internal/base/type"
	roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"
	securityconstant "github.com/ladmakhi81/learning-management-system/internal/security/constant"
	securitytype "github.com/ladmakhi81/learning-management-system/internal/security/type"
)

func (m Middleware) CheckPermissions(allowedPermissions roleentity.Permissions) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authData, authDataExist := ctx.Get(securityconstant.SECURITY_DECODE_KEY)
		if !authDataExist {
			ctx.AbortWithStatusJSON(
				http.StatusForbidden,
				basetype.NewForbiddenAccessResponse(),
			)
			return
		}
		session := authData.(*securitytype.SessionDTO)
		if len(session.Permissions) == 0 {
			ctx.AbortWithStatusJSON(
				http.StatusForbidden,
				basetype.NewForbiddenAccessResponse(),
			)
			return
		}
		hasAccess := true
		mappedUserPermission := make(map[roleentity.Permission]any)
		for _, allowedPermission := range allowedPermissions {
			mappedUserPermission[allowedPermission] = nil
		}
		for _, permission := range session.Permissions {
			if permission == roleentity.SUPER_ADMIN {
				hasAccess = true
				break
			}
			if _, exist := mappedUserPermission[permission]; !exist {
				hasAccess = false
				break
			}
		}
		if !hasAccess {
			ctx.AbortWithStatusJSON(
				http.StatusForbidden,
				basetype.NewForbiddenAccessResponse(),
			)
			return
		}
		ctx.Next()
	}
}
