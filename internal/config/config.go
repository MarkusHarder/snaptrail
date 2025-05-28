package config

import (
	"reflect"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"
)

type Config struct {
	DbUrl         string `env:"DATABASE_URL,required"`
	Dev           bool   `env:"DEV" envDefault:"true"`
	IsDebug       bool   `env:"DEBUG" envdefault:"true"`
	Port          int    `env:"PORT" envDefault:"8115"`
	AdminPort     string `env:"PORT" envDefault:":8000"`
	UiDir         string `env:"UI_DIR" envDefault:"/tmp"`
	DomainSuffix  string `env:"DOMAIN_SUFFIX,required"`
	JwtSecret     string `env:"JWT_SECRET,required"`
	AdminUsername string `env:"ADMIN_USERNAME,required"`
	AdminPassword string `env:"ADMIN_PASSWORD,required"`
	S3Endpoint    string `env:"S3_ENDPOINT,required"`
	S3Bucket      string `env:"S3_BUCKET,required"`
	S3BasePath    string `env:"S3_BASE_PATH"`
	S3Region      string `env:"S3_REGION"`
	S3AccessKey   string `env:"S3_ACCESS_KEY,required"`
	S3SecretKey   string `env:"S3_SECRET_KEY,required"`
}

var config Config

func init() {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal().Msg("Could not parse config")
	}
	trimStringFields(&cfg)
	config = cfg
}

func Get() Config {
	return config
}

func trimStringFields(cfg *Config) {
	val := reflect.ValueOf(cfg).Elem()
	for i := range val.NumField() {
		field := val.Field(i)
		if field.Kind() == reflect.String && field.CanSet() {
			trimmed := strings.TrimSpace(field.String())
			field.SetString(trimmed)
		}
	}
}
