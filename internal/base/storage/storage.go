package basestorage

import (
	"fmt"

	baseconfig "github.com/ladmakhi81/learning-management-system/internal/base/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB     *gorm.DB
	config *baseconfig.Config
}

func NewStorage(config *baseconfig.Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Connect() error {
	connectionString := s.getConnectionString()
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	s.DB = db

	return nil
}

func (s Storage) getConnectionString() string {
	dbHost := s.config.DatabaseConfig.DBHost
	dbUser := s.config.DatabaseConfig.DBUserName
	dbPassword := s.config.DatabaseConfig.DBPassword
	dbName := s.config.DatabaseConfig.DBName
	dbPort := s.config.DatabaseConfig.DBPort

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)
}
