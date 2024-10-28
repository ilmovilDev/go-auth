package env

import (
	"fmt"

	"github.com/ilmovilDev/go-auth/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	AppPort    int    `envconfig:"APP_PORT" required:"true"`
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     int    `envconfig:"DB_PORT" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`
	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
}

func LoadConfig() (*EnvConfig, error) {

	log := logger.NewLogger("Environment")

	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found, loading system environment variables")
	}

	var envs EnvConfig
	err := envconfig.Process("", &envs)

	if err != nil {
		return nil, fmt.Errorf("failed to load envs: %w", err)
	}

	log.Info("Environment Variables initialized successfully")

	return &envs, nil

}
