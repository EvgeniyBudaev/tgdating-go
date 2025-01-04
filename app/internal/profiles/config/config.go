package config

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	AllowOrigins         string `envconfig:"ALLOW_ORIGINS"`
	GatewayPort          string `envconfig:"GATEWAY_PORT"`
	ProfilesHost         string `envconfig:"PROFILES_HOST"`
	Domain               string `envconfig:"DOMAIN"`
	LoggerLevel          string `envconfig:"LOGGER_LEVEL"`
	DBHost               string `envconfig:"POSTGRES_HOST"`
	DBPort               string `envconfig:"POSTGRES_PORT"`
	DBUser               string `envconfig:"POSTGRES_USER"`
	DBPassword           string `envconfig:"POSTGRES_PASSWORD"`
	DBName               string `envconfig:"POSTGRES_NAME"`
	DBSSlMode            string `envconfig:"POSTGRES_SSLMODE"`
	DBSchema             string `envconfig:"POSTGRES_SCHEMA"`
	TelegramBotToken     string `envconfig:"TELEGRAM_BOT_TOKEN"`
	S3AccessKey          string `envconfig:"S3_ACCESS_KEY"`
	S3SecretKey          string `envconfig:"S3_SECRET_KEY"`
	S3EndpointUrl        string `envconfig:"S3_ENDPOINT_URL"`
	S3BucketName         string `envconfig:"S3_BUCKET_NAME"`
	S3BucketPublicDomain string `envconfig:"S3_BUCKET_PUBLIC_DOMAIN"`
	CryptoSecretKey      string `envconfig:"CRYPTO_SECRET_KEY"`
}

func Load(l logger.Logger) (*Config, error) {
	var cfg Config
	err := envconfig.Process("TGDATING", &cfg)
	if err != nil {
		l.Debug("error func Load, method Process by path internal/profiles/config/config.go", zap.Error(err))
		return nil, err
	}
	return &cfg, nil
}
