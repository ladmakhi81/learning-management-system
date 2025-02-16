package baseconfig

import "github.com/spf13/viper"

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
	return &Config{
		DatabaseConfig: newDatabaseConfig(),
	}
}

func newDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Name:     viper.GetString("DB_NAME"),
		Password: viper.GetString("DB_PASSWORD"),
		UserName: viper.GetString("DB_USER"),
		Port:     viper.GetUint("DB_PORT"),
		Host:     viper.GetString("DB_HOST"),
	}
}
