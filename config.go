package main

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	cfg = viper.New()
	lg  = logrus.New()
)

func setupConfig() {
	cfg.AddConfigPath(".")
	cfg.AddConfigPath("$HOME/kgo")
	cfg.AddConfigPath("/etc/kgo")

	cfg.SetConfigName("kgo")
	cfg.SetEnvPrefix("KGO")

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
}
