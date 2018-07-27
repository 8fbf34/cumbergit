package web

var serviceConfig *Config

type Config struct {
	Source string
}

func NewConfig() Config {
	return Config{}
}