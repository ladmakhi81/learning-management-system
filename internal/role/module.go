package role

import (
	"fmt"

	basestorage "github.com/ladmakhi81/learning-management-system/internal/base/storage"
	rolecontractor "github.com/ladmakhi81/learning-management-system/internal/role/contractor"
	rolehandler "github.com/ladmakhi81/learning-management-system/internal/role/handler"
	rolemapper "github.com/ladmakhi81/learning-management-system/internal/role/mapper"
	rolerepository "github.com/ladmakhi81/learning-management-system/internal/role/repository"
	rolerouter "github.com/ladmakhi81/learning-management-system/internal/role/router"
	roleservice "github.com/ladmakhi81/learning-management-system/internal/role/service"
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
	m.container.Provide(rolemapper.NewRoleMapper)
	m.container.Provide(func(storage *basestorage.Storage) rolecontractor.RoleRepository {
		return rolerepository.NewRoleRepositoryImpl(storage)
	})
	m.container.Provide(func(roleRepo rolecontractor.RoleRepository) rolecontractor.RoleService {
		return roleservice.NewRoleServiceImpl(roleRepo)
	})
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
