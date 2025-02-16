package basestorage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func NewStorage() *Storage {
	return &Storage{}
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
	dbHost := "localhost"
	dbUser := "postgres"
	dbPassword := "postgres"
	dbName := "lms_db"
	dbPort := 5432

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)
}
