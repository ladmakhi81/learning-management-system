package basebootstrap

import (
	"fmt"

	"github.com/ladmakhi81/learning-management-system/internal/auth"
	baseconfig "github.com/ladmakhi81/learning-management-system/internal/base/config"
	basestorage "github.com/ladmakhi81/learning-management-system/internal/base/storage"
	"github.com/ladmakhi81/learning-management-system/internal/queue"
	"github.com/ladmakhi81/learning-management-system/internal/role"
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
	viper.AutomaticEnv()
	container := dig.New()
	config := baseconfig.NewConfig()
	if err := config.LoadConfig(); err != nil {
		return fmt.Errorf("environment variable is not loaded : %v", err)
	}
	storage := basestorage.NewStorage(config)
	if err := storage.Connect(); err != nil {
		return fmt.Errorf("database not connected : %v", err)
	}
	b.initializeRedis(container, config)
	b.initializeRabbitmq(container, config)
	container.Provide(func() *baseconfig.Config {
		return config
	})
	container.Provide(func() *basestorage.Storage {
		return storage
	})
	b.container = container
	b.config = config
	if err := b.createSuperAdmin(storage); err != nil {
		return err
	}
	return nil
}

func (b Bootstrap) GetContainer() *dig.Container {
	return b.container
}

func (b Bootstrap) GetConfig() *baseconfig.Config {
	return b.config
}

func (b Bootstrap) initializeRedis(container *dig.Container, config *baseconfig.Config) {
	redisClient := pkgredisclient.NewRedisClient(config)
	redisClient.ConnectRedis()
	container.Provide(func() *pkgredisclient.RedisClient {
		return redisClient
	})
}

func (b Bootstrap) initializeRabbitmq(container *dig.Container, config *baseconfig.Config) error {
	rabbitmqClient, rabbitmqClientErr := pkgrabbitmqclient.NewRabbitmqClient(config.RabbitmqClientURL)
	if rabbitmqClientErr != nil {
		return fmt.Errorf("rabbitmq client not connected : %v", rabbitmqClientErr)
	}
	container.Provide(func() *pkgrabbitmqclient.RabbitmqClient {
		return rabbitmqClient
	})
	return nil
}

func (b Bootstrap) LoadModules() {
	queueModule := queue.NewQueueModule(b.container)
	queueModule.LoadModule()

	roleModule := role.NewRoleModule(b.container)
	roleModule.LoadModule()

	userModule := user.NewUserModule(b.container)
	userModule.LoadModule()

	authModule := auth.NewAuthModule(b.container)
	authModule.LoadModule()

}
