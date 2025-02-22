package authrouter

import (
	"github.com/gin-gonic/gin"
	authhandler "github.com/ladmakhi81/learning-management-system/internal/auth/handler"
	basehandler "github.com/ladmakhi81/learning-management-system/internal/base/handler"
)

type AuthRouter struct {
	apiServer *gin.RouterGroup
	handler   authhandler.AuthHandler
}

func NewAuthRouter(
	apiServer *gin.RouterGroup,
	handler authhandler.AuthHandler,
) AuthRouter {
	return AuthRouter{
		apiServer: apiServer,
		handler:   handler,
	}
}

func (r AuthRouter) SetupRoutes() {
	authApi := r.apiServer.Group("/auth")

	authApi.POST("/login", basehandler.BaseHandler(r.handler.Login))
	authApi.POST("/signup", basehandler.BaseHandler(r.handler.Signup))
}
