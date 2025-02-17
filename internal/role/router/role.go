package rolerouter

import (
	"github.com/gin-gonic/gin"
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

	roleApi.POST("/", r.handler.CreateRole)
	roleApi.GET("/", r.handler.GetRoles)
	roleApi.DELETE("/:id", r.handler.DeleteRoleById)
}
