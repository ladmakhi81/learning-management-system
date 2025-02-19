package rolerouter

import (
	"github.com/gin-gonic/gin"
	basehandler "github.com/ladmakhi81/learning-management-system/internal/base/handler"
	rolehandler "github.com/ladmakhi81/learning-management-system/internal/role/handler"
)

type RoleRouter struct {
	apiServer *gin.RouterGroup
	handler   rolehandler.RoleHandler
}

func NewRoleRouter(
	apiServer *gin.RouterGroup,
	handler rolehandler.RoleHandler,
) RoleRouter {
	return RoleRouter{
		apiServer: apiServer,
		handler:   handler,
	}
}

func (r RoleRouter) SetupRoutes() {
	roleApi := r.apiServer.Group("/roles")

	roleApi.POST("/", basehandler.BaseHandler(r.handler.CreateRole))
	roleApi.GET("/", basehandler.BaseHandler(r.handler.GetRoles))
	roleApi.DELETE("/:id", basehandler.BaseHandler(r.handler.DeleteRoleById))
}
