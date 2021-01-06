package rocketmq

type ConnectConfig struct {
	NameServers []string `toml:"name_servers" json:"name_servers"`
}

type RocketConfig struct {
	*ConnectConfig
	Producers map[string]*ProducerConfig `toml:"producer" json:"producer"`
	Consumers map[string]*ConsumerConfig `toml:"consumer" json:"consumer"`
}
