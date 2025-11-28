package config

import (
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
}

func LoadConfig() *Config {
	port := getEnv("PORT", "8080")
	dbUrl := getEnv("DATABASE_URL", "")
	jwtSecret := getEnv("JWT_SECRET", "")
	env := getEnv("ENV", "developement")

	once.Do(func() {
		cfg = &Config{
			Port:        port,
			DatabaseURL: dbUrl,
			JWTSecret:   jwtSecret,
			Env:         env,
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
