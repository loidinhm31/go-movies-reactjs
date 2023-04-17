package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config App config struct
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

// ServerConfig Server config struct
type ServerConfig struct {
	AppVersion        string
	Port              string
	Mode              string
	ClientId          string
	ClientSecret      string
	SigningKey        string
	TokenTTL          int64
	CookieName        string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
}

// PostgresConfig Postgresql config
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PgDriver           string
}

// LoadConfig Load config file from given path
func LoadConfig(filename string, envProfile string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)

	if envProfile == "" {
		v.AddConfigPath("./config")
		v.AddConfigPath("../../config")
	} else {
		v.AddConfigPath("/etc/config")
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println(err)
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// ParseConfig Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
