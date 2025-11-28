package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
	Env         string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbUrl := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	env := os.Getenv("ENV")
	if env == "" {
		env = "developement"
	}

	return &Config{
		Port:        port,
		DatabaseURL: dbUrl,
		JWTSecret:   jwtSecret,
		Env:         env,
	}
}
