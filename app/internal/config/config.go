package config

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Port                string `envconfig:"AGGREGATION_PORT"`
	LoggerLevel         string `envconfig:"AGGREGATION_LOGGER_LEVEL"`
	Host                string `envconfig:"HOST"`
	DBPort              string `envconfig:"POSTGRES_PORT"`
	DBUser              string `envconfig:"POSTGRES_USER"`
	DBPassword          string `envconfig:"POSTGRES_PASSWORD"`
	DBName              string `envconfig:"POSTGRES_NAME"`
	DBSSlMode           string `envconfig:"POSTGRES_SSLMODE"`
	JWTSecret           string `envconfig:"AGGREGATION_JWT_SECRET"`
	JWTIssuer           string `envconfig:"AGGREGATION_JWT_ISSUER"`
	JWTAudience         string `envconfig:"AGGREGATION_JWT_AUDIENCE"`
	CookieDomain        string `envconfig:"AGGREGATION_COOKIE_DOMAIN"`
	Domain              string `envconfig:"AGGREGATION_DOMAIN"`
	BaseUrl             string `envconfig:"AGGREGATION_KEYCLOAK_BASE_URL"`
	Realm               string `envconfig:"AGGREGATION_KEYCLOAK_REALM"`
	ClientId            string `envconfig:"AGGREGATION_KEYCLOAK_CLIENT_ID"`
	ClientSecret        string `envconfig:"AGGREGATION_KEYCLOAK_CLIENT_SECRET"`
	RealmRS256PublicKey string `envconfig:"AGGREGATION_KEYCLOAK_REALM_RS256_PUBLIC_KEY"`
	TelegramBotToken    string `envconfig:"TELEGRAM_BOT_TOKEN"`
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
