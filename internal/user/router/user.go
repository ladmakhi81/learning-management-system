package userrouter

import (
	"github.com/gin-gonic/gin"
	basehandler "github.com/ladmakhi81/learning-management-system/internal/base/handler"
	roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"
	securitymiddleware "github.com/ladmakhi81/learning-management-system/internal/security/middleware"
	userhandler "github.com/ladmakhi81/learning-management-system/internal/user/handler"
)

type UserRouter struct {
	apiServer  *gin.RouterGroup
	handler    userhandler.UserHandler
	middleware securitymiddleware.Middleware
}

func NewUserRouter(
	apiServer *gin.RouterGroup,
	handler userhandler.UserHandler,
	middleware securitymiddleware.Middleware,
) UserRouter {
	return UserRouter{
		apiServer:  apiServer,
		handler:    handler,
		middleware: middleware,
	}
}

func (r UserRouter) SetupRoutes() {
	userApi := r.apiServer.Group("/users")
	userApi.Use(r.middleware.CheckAccessToken)

	// CREATE USER
	userApi.
		Use(
			r.middleware.CheckPermissions(roleentity.Permissions{
				roleentity.CREATE_USER,
			}),
		).
		POST(
			"/",
			basehandler.BaseHandler(
				r.handler.CreateUser,
			),
		)

	// READ USERS
	userApi.
		Use(
			r.middleware.CheckPermissions(
				roleentity.Permissions{
					roleentity.READ_USER,
				},
			),
		).
		GET(
			"/",
			basehandler.BaseHandler(
				r.handler.GetUsers,
			),
		)

	// ASSIGN ROLE
	userApi.
		Use(
			r.middleware.CheckPermissions(
				roleentity.Permissions{
					roleentity.ASSIGN_ROLE,
				},
			),
		).
		PATCH(
			"/role",
			basehandler.BaseHandler(
				r.handler.AssignRole,
			),
		)

	// UPLOAD PROFILE BY USER
	userApi.POST(
		"/upload-profile",
		basehandler.BaseHandler(
			r.handler.UploadProfileImage,
		),
	)

	// UPLOAD RESUME BY USER
	userApi.POST(
		"/teacher/upload-resume",
		basehandler.BaseHandler(
			r.handler.UploadTeacherResume,
		),
	)

	// UPDATE BASE INFORMATION BY USER
	userApi.PATCH(
		"/base-info",
		basehandler.BaseHandler(
			r.handler.UpdateBaseInformation,
		),
	)

	// CHANGE PASSWORD
	userApi.
		Use(
			r.middleware.CheckPermissions(
				roleentity.Permissions{
					roleentity.CHANGE_PASSWORD,
				},
			),
		).
		PATCH(
			"/change-password",
			basehandler.BaseHandler(
				r.handler.ChangePassword,
			),
		)

	// BLOCK USER
	userApi.
		Use(
			r.middleware.CheckPermissions(
				roleentity.Permissions{
					roleentity.BLOCK_USER,
				},
			),
		).
		PATCH(
			"/block",
			basehandler.BaseHandler(
				r.handler.BlockUser,
			),
		)

	// UNBLOCK USER
	userApi.
		Use(
			r.middleware.CheckPermissions(
				roleentity.Permissions{
					roleentity.BLOCK_USER,
				},
			),
		).
		PATCH(
			"/unblock",
			basehandler.BaseHandler(
				r.handler.UnBlockUser,
			),
		)

	// VERIFY TEACHER
	userApi.
		Use(
			r.middleware.CheckPermissions(
				roleentity.Permissions{
					roleentity.VERIFY_TEACHER,
				},
			),
		).
		PATCH(
			"/teacher/verify",
			basehandler.BaseHandler(
				r.handler.VerifyTeacherByAdmin,
			),
		)

	// COMPLETE USER INFORMATION AS TEACHER
	userApi.PATCH(
		"/teacher/:teacher-id",
		basehandler.BaseHandler(
			r.handler.CompleteTeacherProfile,
		),
	)
}
