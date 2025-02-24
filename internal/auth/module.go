package auth

import (
	"fmt"

	authcontractor "github.com/ladmakhi81/learning-management-system/internal/auth/contractor"
	authhandler "github.com/ladmakhi81/learning-management-system/internal/auth/handler"
	authrouter "github.com/ladmakhi81/learning-management-system/internal/auth/router"
	authservice "github.com/ladmakhi81/learning-management-system/internal/auth/service"
	securitycontractor "github.com/ladmakhi81/learning-management-system/internal/security/contractor"
	usercontractor "github.com/ladmakhi81/learning-management-system/internal/user/contractor"
	pkgredisclient "github.com/ladmakhi81/learning-management-system/pkg/redis-client"
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
	m.container.Provide(func(redisClient *pkgredisclient.RedisClient) authcontractor.SessionService {
		return authservice.NewSessionServiceImpl(redisClient)
	})
	m.container.Provide(func(
		userSvc usercontractor.UserService,
		tokenSvc securitycontractor.TokenService,
		sessionSvc authcontractor.SessionService,
	) authcontractor.AuthService {
		return authservice.NewAuthServiceImpl(userSvc, tokenSvc, sessionSvc)
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
