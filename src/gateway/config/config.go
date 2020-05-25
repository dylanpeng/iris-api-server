package config

import (
	"github.com/BurntSushi/toml"
	"juggernaut/common/env"
	jGrpc "juggernaut/lib/grpc"
	"juggernaut/lib/http"
	"juggernaut/lib/logger"
)

var config *Config

type Config struct {
	Env    *env.Config
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
	config = &Config{}

	if _, err := toml.DecodeFile(file, config); err != nil {
		return err
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
