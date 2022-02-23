package config

import "time"

type REST struct {
	Listen       string        `env:"REST_LISTEN" envDefault:":8080"`
	ReadTimeout  time.Duration `env:"REST_READ_TIMEOUT" envDefault:"30s"`
	WriteTimeout time.Duration `env:"REST_WRITE_TIMEOUT" envDefault:"30s"`
}
