package config

import (
	"juggernaut/common/env"
	jGrpc "juggernaut/lib/grpc"
	"juggernaut/lib/http"
)

type Config struct {
	Env    *env.Config
	Server *ServerConfig `toml:"server"`
}

type ServerConfig struct {
	NetworkInterface string        `toml:"network_interface"`
	BindInterface    bool          `toml:"bind_interface"`
	UseInterfaceIp   bool          `toml:"use_interface_ip"`
	Grpc             *jGrpc.Config `toml:"grpc"`
	Http             *http.Config  `toml:"http"`
}
