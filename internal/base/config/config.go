package baseconfig

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	*DatabaseConfig
}

type DatabaseConfig struct {
	Name     string
	Password string
	UserName string
	Port     uint
	Host     string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) setupConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	viper.AutomaticEnv()

	return nil
}

func (c *Config) loadDatabaseConfig() {
	c.DatabaseConfig = &DatabaseConfig{
		Name:     viper.GetString("DB_NAME"),
		Password: viper.GetString("DB_PASSWORD"),
		UserName: viper.GetString("DB_USER"),
		Port:     viper.GetUint("DB_PORT"),
		Host:     viper.GetString("DB_HOST"),
	}
}

func (c *Config) LoadConfig() error {
	if err := c.setupConfig(); err != nil {

		return err
	}

	c.loadDatabaseConfig()

	return nil
}
