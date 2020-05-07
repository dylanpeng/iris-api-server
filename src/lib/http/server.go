package http

import "time"

type Config struct {
	Host      string    `toml:"host"`
	Port      int       `toml:"port"`
	Charset   string    `toml:"charset"`
	Gzip      bool      `toml:"gzip"`
	PProf     bool      `toml:"pprof"`
	Websocket *WsConfig `toml:"websocket"`
}

type WsConfig struct {
	Enable   bool          `toml:"enable"`
	Endpoint string        `toml:"endpoint"`
	Library  string        `toml:"library"`
	IdleTime time.Duration `toml:"idle_time"`
}
