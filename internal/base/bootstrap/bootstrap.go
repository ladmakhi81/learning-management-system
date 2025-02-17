package basebootstrap

import (
	"fmt"

	baseconfig "github.com/ladmakhi81/learning-management-system/internal/base/config"
	basestorage "github.com/ladmakhi81/learning-management-system/internal/base/storage"
	"github.com/ladmakhi81/learning-management-system/internal/role"
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

	container.Provide(func() *baseconfig.Config {
		return config
	})

	container.Provide(func() *basestorage.Storage {
		return storage
	})

	b.container = container
	b.config = config

	return nil
}

func (b Bootstrap) GetContainer() *dig.Container {
	return b.container
}

func (b Bootstrap) GetConfig() *baseconfig.Config {
	return b.config
}

func (b Bootstrap) LoadModules() {
	roleModule := role.NewRoleModule(b.container)
	roleModule.LoadModule()
}
