package user

import (
	"fmt"

	baseconfig "github.com/ladmakhi81/learning-management-system/internal/base/config"
	basestorage "github.com/ladmakhi81/learning-management-system/internal/base/storage"
	queueservice "github.com/ladmakhi81/learning-management-system/internal/queue/service"
	rolecontractor "github.com/ladmakhi81/learning-management-system/internal/role/contractor"
	usercontractor "github.com/ladmakhi81/learning-management-system/internal/user/contractor"
	userhandler "github.com/ladmakhi81/learning-management-system/internal/user/handler"
	usermapper "github.com/ladmakhi81/learning-management-system/internal/user/mapper"
	userrepository "github.com/ladmakhi81/learning-management-system/internal/user/repository"
	userrouter "github.com/ladmakhi81/learning-management-system/internal/user/router"
	userservice "github.com/ladmakhi81/learning-management-system/internal/user/service"
	"go.uber.org/dig"
)

type UserModule struct {
	container *dig.Container
}

func NewUserModule(
	container *dig.Container,
) UserModule {
	return UserModule{
		container: container,
	}
}

func (m UserModule) LoadModule() {
	m.registerDependencies()
	m.loadRoutes()
}

func (m UserModule) registerDependencies() {
	m.container.Provide(userrouter.NewUserRouter)
	m.container.Provide(userhandler.NewUserHandler)
	m.container.Provide(usermapper.NewUserMapper)
	m.container.Provide(func(storage *basestorage.Storage) usercontractor.UserRepository {
		return userrepository.NewUserRepositoryImpl(storage)
	})
	m.container.Provide(func(
		roleRepo usercontractor.UserRepository,
		config *baseconfig.Config,
		roleSvc rolecontractor.RoleService,
		pdfQueueService *queueservice.PDFQueueService,
	) usercontractor.UserService {
		return userservice.NewUserServiceImpl(roleRepo, config, roleSvc, pdfQueueService)
	})
}

func (m UserModule) loadRoutes() {
	err := m.container.Invoke(func(router userrouter.UserRouter) {
		router.SetupRoutes()
	})

	if err == nil {
		fmt.Println("------ User Module Loaded ------")
	} else {
		fmt.Println("------ User Module Not Loaded : Failed ------")
	}
}
