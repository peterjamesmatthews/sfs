package app

type Config struct {
	JWT_Issuer string `env:"JWT_ISSUER"`
	JWT_Secret []byte `env:"JWT_SECRET"`
}
