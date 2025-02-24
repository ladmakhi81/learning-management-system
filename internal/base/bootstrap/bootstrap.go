package basebootstrap

import (
	"fmt"

	"github.com/ladmakhi81/learning-management-system/internal/auth"
	baseconfig "github.com/ladmakhi81/learning-management-system/internal/base/config"
	basestorage "github.com/ladmakhi81/learning-management-system/internal/base/storage"
	"github.com/ladmakhi81/learning-management-system/internal/queue"
	"github.com/ladmakhi81/learning-management-system/internal/role"
	"github.com/ladmakhi81/learning-management-system/internal/security"
	"github.com/ladmakhi81/learning-management-system/internal/user"
	pkgrabbitmqclient "github.com/ladmakhi81/learning-management-system/pkg/rabbitmq"
	pkgredisclient "github.com/ladmakhi81/learning-management-system/pkg/redis-client"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

type Bootstrap struct {
	container *dig.Container
	config    *baseconfig.Config
}

func NewBootstrap() Bootstrap {
	return Bootstrap{}
}

func (b *Bootstrap) Apply() error {
	container := dig.New()
	b.container = container
	if err := b.initializeConfig(container); err != nil {
		return err
	}
	b.initializeDatabase(container)
	b.initializeRedis(container)
	b.initializeRabbitmq(container)
	return nil
}

func (b Bootstrap) initializeDatabase(container *dig.Container) error {
	storage := basestorage.NewStorage(b.config)
	if err := storage.Connect(); err != nil {
		return fmt.Errorf("database not connected : %v", err)
	}
	container.Provide(func() *basestorage.Storage {
		return storage
	})
	if err := b.createSuperAdmin(storage); err != nil {
		return err
	}
	return nil
}

func (b Bootstrap) initializeRedis(container *dig.Container) {
	redisClient := pkgredisclient.NewRedisClient(b.config)
	redisClient.ConnectRedis()
	container.Provide(func() *pkgredisclient.RedisClient {
		return redisClient
	})
}

func (b Bootstrap) initializeRabbitmq(container *dig.Container) error {
	rabbitmqClient, rabbitmqClientErr := pkgrabbitmqclient.NewRabbitmqClient(b.config.RabbitmqClientURL)
	if rabbitmqClientErr != nil {
		return fmt.Errorf("rabbitmq client not connected : %v", rabbitmqClientErr)
	}
	container.Provide(func() *pkgrabbitmqclient.RabbitmqClient {
		return rabbitmqClient
	})
	return nil
}

func (b *Bootstrap) initializeConfig(container *dig.Container) error {
	viper.AutomaticEnv()
	config := baseconfig.NewConfig()
	b.config = config
	container.Provide(func() *baseconfig.Config {
		return config
	})
	if err := config.LoadConfig(); err != nil {
		return fmt.Errorf("environment variable is not loaded : %v", err)
	}
	return nil
}

func (b Bootstrap) GetContainer() *dig.Container {
	return b.container
}

func (b Bootstrap) GetConfig() *baseconfig.Config {
	return b.config
}

func (b Bootstrap) LoadModules() {
	securityModule := security.NewSecurityModule(b.container)
	securityModule.LoadModule()

	queueModule := queue.NewQueueModule(b.container)
	queueModule.LoadModule()

	roleModule := role.NewRoleModule(b.container)
	roleModule.LoadModule()

	userModule := user.NewUserModule(b.container)
	userModule.LoadModule()

	authModule := auth.NewAuthModule(b.container)
	authModule.LoadModule()
}
