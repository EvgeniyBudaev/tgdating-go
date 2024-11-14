package config

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/logger"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	GatewayPort     string `envconfig:"GATEWAY_PORT"`
	ProfilesPort    string `envconfig:"PROFILES_PORT"`
	Domain          string `envconfig:"DOMAIN"`
	LoggerLevel     string `envconfig:"AGGREGATION_LOGGER_LEVEL"`
	CryptoSecretKey string `envconfig:"CRYPTO_SECRET_KEY"`
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
