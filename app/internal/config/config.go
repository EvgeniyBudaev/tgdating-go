package config

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Port             string `envconfig:"AGGREGATION_PORT"`
	LoggerLevel      string `envconfig:"AGGREGATION_LOGGER_LEVEL"`
	Host             string `envconfig:"HOST"`
	DBPort           string `envconfig:"POSTGRES_PORT"`
	DBUser           string `envconfig:"POSTGRES_USER"`
	DBPassword       string `envconfig:"POSTGRES_PASSWORD"`
	DBName           string `envconfig:"POSTGRES_NAME"`
	DBSSlMode        string `envconfig:"POSTGRES_SSLMODE"`
	TelegramBotToken string `envconfig:"TELEGRAM_BOT_TOKEN"`
}

func Load(l logger.Logger) (*Config, error) {
	var cfg Config
	err := envconfig.Process("TGDATING", &cfg)
	if err != nil {
		l.Debug("error func Load, method Process by path internal/config/config.go", zap.Error(err))
		return nil, err
	}
	return &cfg, nil
}
