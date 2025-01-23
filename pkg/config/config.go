package config

import (
	"log/slog"
	"os"
	"strings"

	"go-gui/pkg/models"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var Settings *models.Config

func Init() {
	if err := parseArgs(); err != nil {
		slog.Error(err.Error())
	}

	viper.SetDefault("log.level", "ERROR")
	viper.SetDefault("config", "./configs/config.yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	viper.SetConfigFile(viper.GetString("config"))

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("Failed to read config file", "error", err)
	}

	if err := viper.Unmarshal(&Settings); err != nil {
		slog.Error("Failed to unmarshal configuration", "error", err)
		os.Exit(1)
	}
}

func parseArgs() error {
	pflag.String("config", "./configs/config.yaml", "Path to configuration file")
	pflag.String("log.level", "", "log level app")
	pflag.Parse()

	return viper.BindPFlags(pflag.CommandLine)
}
