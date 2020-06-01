package env

import "juggernaut/common/coder"

type Config struct {
	Name       string  `toml:"name"`
	Debug      bool    `toml:"debug"`
	HttpCode   string  `toml:"http_code"`
	WsCode     string  `toml:"ws_code"`
	BsCode     string  `toml:"bs_code"`
	DefaultLng float64 `toml:"default_lng"`
	DefaultLat float64 `toml:"default_lat"`
}

var (
	Name       string
	Debug      bool
	HttpCoder  coder.ICoder
	WsCoder    coder.ICoder
	BsCoder    coder.ICoder
	DefaultLng float64
	DefaultLat float64
	GrpcAddr   string
	HttpAddr   string
)

func Init(c *Config) {
	if c.Debug {
		Debug = true
	}

	if c.HttpCode == coder.EncodingProtobuf {
		HttpCoder = coder.ProtoCoder
	} else {
		HttpCoder = coder.JsonCoder
	}

	if c.WsCode == coder.EncodingProtobuf {
		WsCoder = coder.ProtoCoder
	} else {
		WsCoder = coder.JsonCoder
	}

	if c.BsCode == coder.EncodingProtobuf {
		BsCoder = coder.ProtoCoder
	} else {
		BsCoder = coder.JsonCoder
	}

	Name = c.Name
	DefaultLng = c.DefaultLng
	DefaultLat = c.DefaultLat
}
