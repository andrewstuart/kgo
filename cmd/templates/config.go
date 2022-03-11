package main

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	lg = logrus.New()
)

func setupConfig() *viper.Viper {
	cfg := viper.New()
	cfg.AddConfigPath(".")
	cfg.AddConfigPath("$HOME/{{ .Name }}")
	cfg.AddConfigPath("/etc/{{ .Name }}")

	cfg.SetConfigName("{{ .Name }}")
	cfg.SetEnvPrefix("{{ .Name | upper }}")

	cfg.SetDefault("name", "{{ .Name | upper }}")

	cfg.AutomaticEnv()
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := cfg.ReadInConfig(); err != nil {
		lg.WithError(err).Error("could not read initial config")
	}

	cfg.OnConfigChange(func(_ fsnotify.Event) {
		if err := cfg.ReadInConfig(); err != nil {
			lg.WithError(err).Warn("could not reload config")
		}
	})

	go cfg.WatchConfig()

	return cfg
}
