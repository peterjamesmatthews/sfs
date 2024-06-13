package server

type Config struct {
	Hostname      string `env:"HOSTNAME"`
	Port          int    `env:"PORT"`
	GraphEndpoint string `env:"GRAPH_ENDPOINT"`
}
