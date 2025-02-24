package security

import (
	"fmt"

	baseconfig "github.com/ladmakhi81/learning-management-system/internal/base/config"
	securitycontractor "github.com/ladmakhi81/learning-management-system/internal/security/contractor"
	securitymiddleware "github.com/ladmakhi81/learning-management-system/internal/security/middleware"
	securityservice "github.com/ladmakhi81/learning-management-system/internal/security/service"
	pkgredisclient "github.com/ladmakhi81/learning-management-system/pkg/redis-client"
	"go.uber.org/dig"
)

type SecurityModule struct {
	container *dig.Container
}

func NewSecurityModule(
	container *dig.Container,
) SecurityModule {
	return SecurityModule{
		container: container,
	}
}

func (m SecurityModule) LoadModule() {
	m.registerDependencies()

	fmt.Println("------ Security Module Load ------")
}

func (m SecurityModule) registerDependencies() {
	m.container.Provide(func(config *baseconfig.Config) securitycontractor.TokenService {
		return securityservice.NewTokenServiceImpl(config)
	})
	m.container.Provide(func(redisClient *pkgredisclient.RedisClient) securitycontractor.SessionService {
		return securityservice.NewSessionServiceImpl(redisClient)
	})
	m.container.Provide(securitymiddleware.NewMiddleware)
}
