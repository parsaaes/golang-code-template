package config

import (
	"bytes"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type (
	Config struct {
		Server     Server     `mapstructure:"server" validate:"required"`
		Logger     Logger     `mapstructure:"logger" validate:"required"`
		Redis      Redis      `mapstructure:"redis" validate:"required"`
		Postgres   Postgres   `mapstructure:"postgres" validate:"required"`
		Monitoring Monitoring `mapstructure:"monitoring" validate:"required"`
	}

	Server struct {
		ReadTimeout     time.Duration `mapstructure:"read-timeout" validate:"required"`
		WriteTimeout    time.Duration `mapstructure:"write-timeout" validate:"required"`
		GracefulTimeout time.Duration `mapstructure:"graceful-timeout" validate:"required"`
		Address         string        `mapstructure:"address" validate:"required"`
	}

	Logger struct {
		Level string `mapstructure:"level" validate:"required"`
	}

	Postgres struct {
		Host               string        `mapstructure:"host" validate:"required"`
		Port               int           `mapstructure:"port" validate:"required"`
		Username           string        `mapstructure:"user" validate:"required"`
		Password           string        `mapstructure:"pass" validate:"required"`
		DBName             string        `mapstructure:"dbname" validate:"required"`
		ConnectTimeout     time.Duration `mapstructure:"connect-timeout" validate:"required"`
		ConnectionLifetime time.Duration `mapstructure:"connection-lifetime" validate:"required"`
		MaxOpenConnections int           `mapstructure:"max-open-connections" validate:"required"`
		MaxIdleConnections int           `mapstructure:"max-idle-connections" validate:"required"`
	}

	Redis struct {
		Master RedisConfig `mapstructure:"master" validate:"required"`
		Slave  RedisConfig `mapstructure:"slave" validate:"required"`
	}

	RedisConfig struct {
		Address         string        `mapstructure:"address" validate:"required"`
		PoolSize        int           `mapstructure:"pool-size"`
		MinIdleConns    int           `mapstructure:"min-idle-conns"`
		DialTimeout     time.Duration `mapstructure:"dial-timeout"`
		ReadTimeout     time.Duration `mapstructure:"read-timeout"`
		WriteTimeout    time.Duration `mapstructure:"write-timeout"`
		PoolTimeout     time.Duration `mapstructure:"pool-timeout"`
		IdleTimeout     time.Duration `mapstructure:"idle-timeout"`
		MaxRetries      int           `mapstructure:"max-retries"`
		MinRetryBackoff time.Duration `mapstructure:"min-retry-backoff"`
		MaxRetryBackoff time.Duration `mapstructure:"max-retry-backoff"`
	}

	Prometheus struct {
		Enabled bool   `mapstructure:"enabled"`
		Address string `mapstructure:"address" validate:"required"`
	}

	Monitoring struct {
		Prometheus Prometheus `mapstructure:"prometheus" validate:"required"`
	}
)

func (c Config) Validate() error {
	return validator.New().Struct(c)
}

func Init(path string) Config {
	var cfg Config

	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewReader([]byte(Default))); err != nil {
		logrus.Panicf("error loading default configs: %s", err.Error())
	}

	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	v.SetEnvPrefix(Namespace)
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	err := v.MergeInConfig()
	if err != nil {
		logrus.Warn("no config file found. Using defaults and environment variables")
	}

	if err := v.UnmarshalExact(&cfg); err != nil {
		logrus.Fatalf("invalid configuration: %s", err)
	}

	if err := cfg.Validate(); err != nil {
		logrus.Fatalf("invalid configuration: %s", err)
	}

	return cfg
}
