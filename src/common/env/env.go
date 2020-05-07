package env

type Config struct {
	Name       string  `toml:"name"`
	Debug      bool    `toml:"debug"`
	HttpCode   string  `toml:"http_code"`
	WsCode     string  `toml:"ws_code"`
	BsCode     string  `toml:"bs_code"`
	DefaultLng float64 `toml:"default_lng"`
	DefaultLat float64 `toml:"default_lat"`
}
