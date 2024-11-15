package config

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	LoggerLevel      string `envconfig:"LOGGER_LEVEL"`
	TelegramBotToken string `envconfig:"TELEGRAM_BOT_TOKEN"`
}

func Load(l logger.Logger) (*Config, error) {
	var cfg Config
	err := envconfig.Process("TGDATING", &cfg)
	if err != nil {
		l.Debug("error func Load, method Process by path internal/telegram/config/config.go", zap.Error(err))
		return nil, err
	}
	return &cfg, nil
}
