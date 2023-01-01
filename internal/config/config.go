package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseHost       string        `mapstructure:"DATABASE_HOST"`
	DatabasePort       int           `mapstructure:"DATABASE_PORT"`
	DatabaseName       string        `mapstructure:"DATABASE_NAME"`
	DatabaseUser       string        `mapstructure:"DATABASE_USER"`
	DatabasePass       string        `mapstructure:"DATABASE_PASS"`
	DatabaseSSLMode    string        `mapstructure:"DATABASE_SSL_MODE"`
	Port               int           `mapstructure:"PORT"`
	HTTPMaxHeaderBytes int           `mapstructure:"HTTP_MAX_HEADER_BYTES"`
	HTTPReadTimeout    time.Duration `mapstructure:"HTTP_READ_TIMEOUT"`
	HTTPWriteTimeout   time.Duration `mapstructure:"HTTP_WRITE_TIMEOUT"`
	JWTAccessTokenTTL  string        `mapstructure:"JWT_ACCESS_TOKEN_TTL"`
	JWTRefreshTokenTTL string        `mapstructure:"JWT_REFRESH_TOKEN_TTL"`
}

func Init() (*Config, error) {
	if err := parseConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

var configsPath = "./configs"

func parseConfig() error {
	viper.AddConfigPath(configsPath)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
