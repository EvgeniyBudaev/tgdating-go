package config

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/logger"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	AllowOrigins     string `envconfig:"ALLOW_ORIGINS"`
	GatewayPort      string `envconfig:"GATEWAY_PORT"`
	ProfilesHost     string `envconfig:"PROFILES_HOST"`
	LoggerLevel      string `envconfig:"LOGGER_LEVEL"`
	TelegramBotToken string `envconfig:"TELEGRAM_BOT_TOKEN"`
	CryptoSecretKey  string `envconfig:"CRYPTO_SECRET_KEY"`
	Kafka1           string `envconfig:"KAFKA_1"`
	Kafka2           string `envconfig:"KAFKA_2"`
	Kafka3           string `envconfig:"KAFKA_3"`
}

func Load(l logger.Logger) (*Config, error) {
	var cfg Config
	err := envconfig.Process("TGDATING", &cfg)
	if err != nil {
		l.Debug("error func Load, method Process by path internal/gateway/config/config.go", zap.Error(err))
		return nil, err
	}
	return &cfg, nil
}
