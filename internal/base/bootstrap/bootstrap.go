package basebootstrap

import (
	"log"

	baseconfig "github.com/ladmakhi81/learning-management-system/internal/base/config"
	basestorage "github.com/ladmakhi81/learning-management-system/internal/base/storage"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

type Bootstrap struct{}

func NewBootstrap() Bootstrap {
	return Bootstrap{}
}

func (b Bootstrap) Apply() {
	viper.AutomaticEnv()
	container := dig.New()
	config := baseconfig.NewConfig()

	if err := config.LoadConfig(); err != nil {
		log.Fatalf("environment variable is not loaded : %v", err)
	}

	storage := basestorage.NewStorage(config)

	if err := storage.Connect(); err != nil {
		log.Fatalf("database not connected : %v", err)
	}

	container.Provide(func() *baseconfig.Config {
		return config
	})

	container.Provide(func() *basestorage.Storage {
		return storage
	})
}
