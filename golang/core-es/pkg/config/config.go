package config

import (
	"melody-io/core-es/pkg/env"
	"melody-io/core-es/pkg/log"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"gopkg.in/yaml.v2"
)

const (
	defaultJWTExpirationHours = 72
)

type Config struct {
	RSocketHost   string `yaml:"rsocket_host" env:"APP_RSOCKET_HOST"`
	RSocketPort   int    `yaml:"rsocket_port" env:"APP_RSOCKET_PORT"`
	GatewayServe  string `yaml:"gw_serve_uri" env:"APP_GW_SERVE_URI"`
	JWTSigningKey string `yaml:"jwt_signing_key" env:"APP_JWT_SIGNING_KEY,secret"`
	JWTExpiration int    `yaml:"jwt_expiration" env:"APP_JWT_EXPIRATION"`
	MongoUrl      string `yaml:"MONGO_URL" env:"APP_MONGO_URL"`
	NATSUrl       string `yaml:"NATS_URL" env:"APP_NATS_URL"`
}

// Validate validates the application configuration.
func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.JWTSigningKey, validation.Required),
	)
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string, logger log.Logger) (*Config, error) {
	// default config
	c := Config{
		JWTExpiration: defaultJWTExpirationHours,
	}

	// load from YAML config file
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// load from environment variables prefixed with "APP_"
	if err = env.New("APP_", logger.Infof).Load(&c); err != nil {
		return nil, err
	}

	// validation
	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, err
}
