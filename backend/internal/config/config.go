package config

import (
	"os"
)

type Config struct {
	Port     string
	DB       string
	Issuer   string
	Audience string
}

func FromEnv() Config {
	return Config{
		Port:     get("API_PORT", "8081"),
		DB:       must("DB_DSN"),
		Issuer:   must("JWT_ISSUER"),
		Audience: must("JWT_AUDIENCE"),
	}
}
func must(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic("missing " + k)
	}
	return v
}
func get(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}
