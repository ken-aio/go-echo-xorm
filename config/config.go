package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

// LoadConfig loading config file
func LoadConfig(e *echo.Echo) {
	viper.SetConfigType("yaml")
	confName := "local"
	confPath := "config"
	if env := os.Getenv("echo.env"); env != "" {
		confName = env
	}
	if env := os.Getenv("conf.path"); env != "" {
		confPath = env
	}
	viper.SetConfigName(confName)
	viper.AddConfigPath(confPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %+v", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(ev fsnotify.Event) {
		e.Logger.Info("Config file changed: ", ev.Name)
	})
}
