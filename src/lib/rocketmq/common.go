package rocketmq

type ConnectConfig struct {
	NameServers []string `toml:"name_servers" json:"name_servers"`
}

type Message struct {
	Tag     string
	Payload []byte
	Props   map[string]string
}

type RocketConfig struct {
	*ConnectConfig
	Producers map[string]*ProducerConfig `toml:"producer" json:"producer"`
	Consumers map[string]*ConsumerConfig `toml:"consumer" json:"consumer"`
}
