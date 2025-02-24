package auth

import (
	"fmt"

	authcontractor "github.com/ladmakhi81/learning-management-system/internal/auth/contractor"
	authhandler "github.com/ladmakhi81/learning-management-system/internal/auth/handler"
	authrouter "github.com/ladmakhi81/learning-management-system/internal/auth/router"
	authservice "github.com/ladmakhi81/learning-management-system/internal/auth/service"
	rolecontractor "github.com/ladmakhi81/learning-management-system/internal/role/contractor"
	securitycontractor "github.com/ladmakhi81/learning-management-system/internal/security/contractor"
	usercontractor "github.com/ladmakhi81/learning-management-system/internal/user/contractor"
	"go.uber.org/dig"
)

type AuthModule struct {
	container *dig.Container
}

func NewAuthModule(
	container *dig.Container,
) AuthModule {
	return AuthModule{
		container: container,
	}
}

func (m AuthModule) LoadModule() {
	m.registerDependencies()
	m.loadRoutes()
}

func (m AuthModule) registerDependencies() {
	m.container.Provide(authrouter.NewAuthRouter)
	m.container.Provide(authhandler.NewAuthHandler)
	m.container.Provide(func(
		userSvc usercontractor.UserService,
		tokenSvc securitycontractor.TokenService,
		sessionSvc securitycontractor.SessionService,
		roleSvc rolecontractor.RoleService,
	) authcontractor.AuthService {
		return authservice.NewAuthServiceImpl(userSvc, tokenSvc, sessionSvc, roleSvc)
	})
}

func (m AuthModule) loadRoutes() {
	err := m.container.Invoke(func(router authrouter.AuthRouter) {
		router.SetupRoutes()
	})

	if err == nil {
		fmt.Println("------ Auth Module Loaded ------")
	} else {
		fmt.Println("------ Auth Module Not Loaded : Failed ------", err)
	}
}
