package config

import (
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
)

type Config struct {
	Primary       Primary              `koanf:"primary" validate :"required"`
	Server        ServerConfig         `koanf:"server" validate :"required"`
	Database      DatabaseConfig       `koanf:"database" validate :"required"`
	Auth          AuthConfig           `koanf:"auth" validate :"required"`
	Redis         RedisConfig          `koanf:"redis_config" validate :"required"`
	Observability *ObservabilityConfig `koanf:"observabiltyg"`
}

type Primary struct {
	Env string `koanf:"env" validate :"required"`
}

type ServerConfig struct {
	Port               string   `koanf:"port" validate :"required"`
	ReadTimeout        int      `koanf:"read_timeout" validate :"required"`
	WriteTimeout       int      `koanf:"write_timeout" validate :"required"`
	IdleTImeout        int      `koanf:"idle_timeout" validate :"required"`
	CORSAllowedOrigins []string `koanf:"cors_allowed_origins" validate :"required"`
}

type DatabaseConfig struct {
	Host            string `koanf:"host" validate :"required"`
	Port            int    `koanf:"port" validate :"required"`
	User            string `koanf:"user" validate :"required"`
	Password        string `koanf:"password"`
	Name            string `koanf:"name" validate :"required"`
	SSLMode         string `koanf:"ssl_node" validate :"required"`
	MaxOpenConns    int    `koanf:"max_open_conns" validate :"required"`
	MaxIdleConns    int    `koanf:"max_idle_conns" validate :"required"`
	ConnMaxLifetime int    `koanf:"conn_max_lifetime" validate :"required"`
	ConnMaxIdleTime int    `koanf:"conn_max_idle_time" validate :"required"`
}

type RedisConfig struct {
	Address string `koanf:"address" validate :"required"`
}
type AuthConfig struct {
	SecretKey string `koanf:"secret_key" validate :"required"`
}

func LoadConfig() (*Config, error) {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	k := koanf.New(".")

	err := k.Load(env.Provider("ECHO_", ".", func(s string) string {
		return strings.ToLower(strings.TrimPrefix(s, "ECHO_"))
	}), nil)

	if err != nil {
		logger.Fatal().Err(err).Msg("could not load initial env variable")
	}

	mainConfig := &Config{}

	err = k.Unmarshal("", mainConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("could not unmarshial main config")
	}

	validate := validator.New()

	err = validate.Struct(mainConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("config validation failed")
	}

	//set default observabilty config if provided
	if mainConfig.Observability == nil {
		mainConfig.Observability = DefaultObservabilityConfig()
	}

	//override sercice name and environment from primary config
	mainConfig.Observability.ServiceName = "go-echo"
	mainConfig.Observability.Environment = mainConfig.Primary.Env

	//validate observabilty config
	if err := mainConfig.Observability.Validate(); err != nil {
		logger.Fatal().Err(err).Msg("Invalid observability config")
	}

	return mainConfig, nil
}
