package config

import "os"

type Config struct {
 DatabaseURL string
 Port        string
}

func Load() Config {
 return Config{
  DatabaseURL: os.Getenv("DATABASE_URL"),
  Port:        getEnv("PORT", "8080"),
 }
}

func getEnv(k, d string) string {
 if v := os.Getenv(k); v != "" {
  return v
 }
 return d
}