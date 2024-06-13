package app

type Config struct {
	AUTH0_DOMAIN string `env:"AUTH0_DOMAIN"`
	JWT_ISSUER   string `env:"JWT_ISSUER"`
	JWT_SECRET   []byte `env:"JWT_SECRET"`
}
