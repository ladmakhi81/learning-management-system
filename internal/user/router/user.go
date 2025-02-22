package userrouter

import (
	"github.com/gin-gonic/gin"
	basehandler "github.com/ladmakhi81/learning-management-system/internal/base/handler"
	userhandler "github.com/ladmakhi81/learning-management-system/internal/user/handler"
)

type UserRouter struct {
	apiServer *gin.RouterGroup
	handler   userhandler.UserHandler
}

func NewUserRouter(
	apiServer *gin.RouterGroup,
	handler userhandler.UserHandler,
) UserRouter {
	return UserRouter{
		apiServer: apiServer,
		handler:   handler,
	}
}

func (r UserRouter) SetupRoutes() {
	userApi := r.apiServer.Group("/users")

	userApi.POST("/", basehandler.BaseHandler(r.handler.CreateUser))
	userApi.GET("/", basehandler.BaseHandler(r.handler.GetUsers))
	userApi.POST("/upload-profile", basehandler.BaseHandler(r.handler.UploadProfileImage))
	userApi.POST("/teacher/upload-resume", basehandler.BaseHandler(r.handler.UploadTeacherResume))
	userApi.PATCH("/base-info", basehandler.BaseHandler(r.handler.UpdateBaseInformation))
	userApi.PATCH("/change-password", basehandler.BaseHandler(r.handler.ChangePassword))
	userApi.PATCH("/block", basehandler.BaseHandler(r.handler.BlockUser))
	userApi.PATCH("/unblock", basehandler.BaseHandler(r.handler.UnBlockUser))
	userApi.PATCH("/teacher/verify", basehandler.BaseHandler(r.handler.VerifyTeacherByAdmin))
	userApi.PATCH("/teacher/:teacher-id", basehandler.BaseHandler(r.handler.CompleteTeacherProfile))
}
