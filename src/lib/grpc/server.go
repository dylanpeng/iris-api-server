package grpc

import "strconv"

type Config struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func (c *Config) GetAddr() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}
