package configs

import (
	"log"
	"os"

	"github.com/mvpoyatt/go-api/api"
	"github.com/mvpoyatt/go-api/database"
	"github.com/spf13/viper"
)

var (
	Environment string
	Values      ConfigValues
)

type ConfigValues struct {
	LoggerLevel string            `mapstructure:"logger_level"`
	Database    database.DbConfig `mapstructure:"db"`
	Server      api.ServerConfig  `mapstructure:"server"`
}

const (
	Development = "development"
	Production  = "production"
)

func init() {
	Environment = os.Getenv("ENVIRONMENT")
	if Environment == "" {
		Environment = "development"
	}
}

func LoadConfigs() {
	filePath := "configs/" + Environment + ".yaml"
	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicln("Cannot read config file: %w", err)
	}
	if err := viper.Unmarshal(&Values); err != nil {
		log.Panicln("Cannot load config file: %w", err)
	}
}
