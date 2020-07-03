package config

import (
	"github.com/BurntSushi/toml"
	"juggernaut/common/env"
	"juggernaut/lib/kafka"
	"juggernaut/lib/logger"
)

var config *Config

type Config struct {
	Env      *env.Config                      `toml:"env"`
	Consumer map[string]*kafka.ConsumerConfig `toml:"kafka_consumer"`
	Log      *logger.Config                   `toml:"log"`
}

func Init(file string) error {
	config = &Config{Log: logger.DefaultConfig()}

	if _, err := toml.DecodeFile(file, config); err != nil {
		return err
	}

	env.Init(config.Env)

	if env.Debug {
		config.Log.Level = logger.Levels[logger.DebugLevel].Name
	}

	return nil
}

func GetConfig() *Config {
	return config
}

func GetConsumer(name string) *kafka.ConsumerConfig {
	return config.Consumer[name]
}

func GetLog() *logger.Config {
	return config.Log
}
