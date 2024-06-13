package db

import "fmt"

type Config struct {
	Hostname string `env:"HOSTNAME"`
	Port     int    `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME"`
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=disable",
		"postgres",
		c.User,
		c.Password,
		c.Hostname,
		c.Port,
		c.Name,
	)
}
