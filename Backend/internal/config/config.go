package config

import (
	"log"
	"os"
	"sync"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
	Env         string
	Issuer      string
}

func LoadConfig() *Config {
	port := getEnv("PORT", "8080")
	dbUrl := getEnv("DATABASE_URL", "")
	jwtSecret := getEnv("JWT_SECRET", "")
	env := getEnv("ENV", "developement")
	issuer := getEnv("ISSUER", "default")

	if issuer == "default" {
		log.Fatal("ISSUER environment variable must be set")
	}

	once.Do(func() {
		cfg = &Config{
			Port:        port,
			DatabaseURL: dbUrl,
			JWTSecret:   jwtSecret,
			Env:         env,
			Issuer:      issuer,
		}
	})

	return cfg
}

func Get() *Config {
	return cfg
}

func Set(key string, value string) *Config {
	switch key {
	}
	return cfg
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
