package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"juggernaut/common/env"
	jGrpc "juggernaut/lib/grpc"
	"juggernaut/lib/http"
	"juggernaut/lib/logger"
	"juggernaut/lib/net2"
)

var config *Config

type Config struct {
	Env    *env.Config    `toml:"env"`
	Server *ServerConfig  `toml:"server"`
	Log    *logger.Config `toml:"log"`
}

type ServerConfig struct {
	NetworkInterface string        `toml:"network_interface"`
	BindInterface    bool          `toml:"bind_interface"`
	UseInterfaceIp   bool          `toml:"use_interface_ip"`
	Grpc             *jGrpc.Config `toml:"grpc"`
	Http             *http.Config  `toml:"http"`
}

func Init(file string) error {
	config = &Config{
		Server: &ServerConfig{
			Http: http.DefaultConfig(),
		},
		Log: logger.DefaultConfig(),
	}

	if _, err := toml.DecodeFile(file, config); err != nil {
		return err
	}

	env.Init(config.Env)

	if config.Server.NetworkInterface != "" {
		interfaceIp, err := net2.GetInterfaceIp(config.Server.NetworkInterface)

		if err != nil {
			return err
		}

		if config.Server.BindInterface {
			config.Server.Http.Host = interfaceIp
			config.Server.Grpc.Host = interfaceIp
		}

		if config.Server.UseInterfaceIp {
			env.HttpAddr = fmt.Sprintf("%s:%d", interfaceIp, config.Server.Http.Port)
			env.GrpcAddr = fmt.Sprintf("%s:%d", interfaceIp, config.Server.Grpc.Port)
		}
	}

	if env.Debug {
		config.Log.Level = logger.Levels[logger.DebugLevel].Name
	}

	config.Server.Http.Log = &http.LogConfig{
		Level:      config.Log.Level,
		Color:      config.Log.Color,
		TimeFormat: config.Log.TimeFormat,
	}

	return nil
}

func GetConfig() *Config {
	return config
}

func GetGrpc() *jGrpc.Config {
	return config.Server.Grpc
}

func GetHttp() *http.Config {
	return config.Server.Http
}

func GetLog() *logger.Config {
	return config.Log
}
