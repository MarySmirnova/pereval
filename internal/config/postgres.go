package config

type Postgres struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Database string `env:"DB_DATABASE" envDefault:"pereval"`
	User     string `env:"DB_USER" envDefault:"selectel"`
	Password string `env:"DB_PASSWORD" envDefault:"selectel"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
}
