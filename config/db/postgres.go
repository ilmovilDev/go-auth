package db

import (
	"fmt"

	"github.com/ilmovilDev/go-auth/config/env"
	"github.com/ilmovilDev/go-auth/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres(envs *env.EnvConfig) (*gorm.DB, error) {

	log := logger.NewLogger("Database")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		envs.DBHost, envs.DBUser, envs.DBPassword, envs.DBName, envs.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Info("Database initialized successfully")

	return db, nil

}
