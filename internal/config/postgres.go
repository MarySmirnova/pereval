package config

type Postgres struct {
	Host     string `env:"DB_HOST"`
	Database string `env:"DB_DATABASE"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Port     int    `env:"DB_PORT"`
}
