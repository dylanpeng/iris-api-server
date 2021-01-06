package config

import (
	"github.com/BurntSushi/toml"
	"juggernaut/common/env"
	"juggernaut/lib/kafka"
	"juggernaut/lib/logger"
	"juggernaut/lib/rocketmq"
)

var config *Config

type Config struct {
	Env      *env.Config                      `toml:"env"`
	Consumer map[string]*kafka.ConsumerConfig `toml:"kafka_consumer"`
	Log      *logger.Config                   `toml:"log"`
	Rocket   *rocketmq.RocketConfig           `toml:"rocket"`
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

	if config.Rocket != nil {
		if config.Rocket.Producers != nil {
			for _, v := range config.Rocket.Producers {
				v.ConnectConfig = config.Rocket.ConnectConfig
			}
		}

		if config.Rocket.Consumers != nil {
			for _, v := range config.Rocket.Consumers {
				v.ConnectConfig = config.Rocket.ConnectConfig
			}
		}
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
