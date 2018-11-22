package main

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	cfg = viper.New()
	lg  = logrus.New()
)

func setupConfig() {
	cfg.AddConfigPath(".")
	cfg.AddConfigPath("$HOME/{{ .Name }}")
	cfg.AddConfigPath("/etc/{{ .Name }}")

	cfg.SetConfigName("{{ .Name }}")
	cfg.SetEnvPrefix("{{ .Name | upper }}")

	if err := viper.ReadInConfig(); err != nil {
		lg.WithError(err).Error("could not read initial config")
	}

	cfg.AutomaticEnv()
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	cfg.OnConfigChange(func(_ fsnotify.Event) {
		if err := viper.ReadInConfig(); err != nil {
			lg.WithError(err).Warn("could not reload config")
		}
	})

	cfg.WatchConfig()
}
