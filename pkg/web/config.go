package web

var serviceConfig *Config

type Config struct {
	DefaultRoute  string
	JsonRoute     []string
	YamlRoute     []string
}

func NewConfig() Config {
	return Config{DefaultRoute: "/api/customers"}
}