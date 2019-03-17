package config

import (
	"os"

	"github.com/spf13/viper"
)

// Viper viper
var Config *viper.Viper
var rootDir string

func env() {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.AddConfigPath(rootDir)
	err := Config.ReadInConfig()
	if err != nil {
		panic("read config error")
	}
}

func init() {
	var err error
	rootDir, err = os.Getwd()
	if err != nil {
		panic("init config exception, because os.Getwd() failure")
	}
	env()
}
