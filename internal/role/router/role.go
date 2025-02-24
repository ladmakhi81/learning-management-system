package rolerouter

import (
	"github.com/gin-gonic/gin"
	basehandler "github.com/ladmakhi81/learning-management-system/internal/base/handler"
	roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"
	rolehandler "github.com/ladmakhi81/learning-management-system/internal/role/handler"
	securitymiddleware "github.com/ladmakhi81/learning-management-system/internal/security/middleware"
)

type RoleRouter struct {
	apiServer  *gin.RouterGroup
	handler    rolehandler.RoleHandler
	middleware securitymiddleware.Middleware
}

func NewRoleRouter(
	apiServer *gin.RouterGroup,
	handler rolehandler.RoleHandler,
	middleware securitymiddleware.Middleware,
) RoleRouter {
	return RoleRouter{
		apiServer:  apiServer,
		handler:    handler,
		middleware: middleware,
	}
}

func (r RoleRouter) SetupRoutes() {
	roleApi := r.apiServer.Group("/roles")

	roleApi.Use(
		r.middleware.CheckAccessToken,
		r.middleware.CheckPermissions(
			roleentity.Permissions{
				roleentity.CREATE_ROLE,
				roleentity.DELETE_ROLE,
				roleentity.EDIT_ROLE,
				roleentity.READ_ROLE,
			},
		),
	)

	roleApi.POST(
		"/",
		basehandler.BaseHandler(
			r.handler.CreateRole,
		),
	)

	roleApi.GET("/", basehandler.BaseHandler(r.handler.GetRoles))
	roleApi.DELETE("/:id", basehandler.BaseHandler(r.handler.DeleteRoleById))
}
