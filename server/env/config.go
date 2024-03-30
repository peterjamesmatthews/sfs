package env

type Config struct {
	Server ServerConfig `env:", prefix=SERVER_"`
}

type ServerConfig struct {
	Hostname string `env:"HOSTNAME"`
	Port     string `env:"PORT"`
	Endpoint string `env:"GRAPH_ENDPOINT"`
}
