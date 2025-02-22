package baseconfig

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	*DatabaseConfig
	*ServerConfig
}

type ServerConfig struct {
	ServerPort      uint
	UploadDirectory string
	SecretKey       string
}

type DatabaseConfig struct {
	DBName     string
	DBPassword string
	DBUserName string
	DBPort     uint
	DBHost     string
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

func (c *Config) loadServerConfig() {
	c.ServerConfig = &ServerConfig{
		ServerPort:      viper.GetUint("APP_PORT"),
		UploadDirectory: viper.GetString("UPLOAD_DIR"),
		SecretKey:       viper.GetString("SECRET_KEY"),
	}
}

func (c *Config) loadDatabaseConfig() {
	c.DatabaseConfig = &DatabaseConfig{
		DBName:     viper.GetString("DB_NAME"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBUserName: viper.GetString("DB_USER"),
		DBPort:     viper.GetUint("DB_PORT"),
		DBHost:     viper.GetString("DB_HOST"),
	}
}

func (c *Config) LoadConfig() error {
	if err := c.setupConfig(); err != nil {

		return err
	}

	c.loadDatabaseConfig()
	c.loadServerConfig()

	return nil
}
