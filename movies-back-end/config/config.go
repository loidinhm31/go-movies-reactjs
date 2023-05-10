package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config App config struct
type Config struct {
	Server     ServerConfig
	Postgres   PostgresConfig
	Keycloak   KeycloakConfig
	Cloudinary CloudinaryConfig
}

// ServerConfig Server config struct
type ServerConfig struct {
	AppVersion        string
	Port              string
	Mode              string
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

type KeycloakConfig struct {
	EndPoint     string
	ClientId     string // clientId specified in KeycloakConfig
	ClientSecret string // client secret specified in KeycloakConfig
	Realm        string // realm specified in KeycloakConfig
}

type CloudinaryConfig struct {
	CloudName  string
	ApiKey     string
	ApiSecret  string
	FolderPath string
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
