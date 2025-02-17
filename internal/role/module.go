package role

import (
	"fmt"

	rolehandler "github.com/ladmakhi81/learning-management-system/internal/role/handler"
	rolerouter "github.com/ladmakhi81/learning-management-system/internal/role/router"
	"go.uber.org/dig"
)

type RoleModule struct {
	container *dig.Container
}

func NewRoleModule(
	container *dig.Container,
) RoleModule {
	return RoleModule{
		container: container,
	}
}

func (m RoleModule) LoadModule() {
	m.registerDependencies()
	m.loadRoutes()
}

func (m RoleModule) registerDependencies() {
	m.container.Provide(rolerouter.NewRoleRouter)
	m.container.Provide(rolehandler.NewRoleHandler)
}

func (m RoleModule) loadRoutes() {
	err := m.container.Invoke(
		func(router rolerouter.RoleRouter) {
			router.SetupRoutes()
		},
	)

	if err == nil {
		fmt.Println("role module loaded successfully")
	} else {
		fmt.Println("role module failed to load", err)
	}
}
