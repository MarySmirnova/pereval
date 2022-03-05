package config

type Postgres struct {
	Host     string `env:"FSTR_DB_HOST"`
	Database string `env:"FSTR_DB_DATABASE"`
	User     string `env:"FSTR_DB_LOGIN"`
	Password string `env:"FSTR_DB_PASS"`
	Port     int    `env:"FSTR_DB_PORT"`
}
