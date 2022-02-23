package config

type Application struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"DEBUG"`

	REST     REST
	Postgres Postgres
}
