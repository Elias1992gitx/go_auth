package config

import "os"

type Config struct {
    MongoURI    string
    JWTSecret   string
    Port        string
    AllowOrigins string
}

func LoadConfig() *Config {
    return &Config{
        MongoURI:     getEnvOrDefault("MONGODB_URI", ""),
        JWTSecret:    getEnvOrDefault("JWT_SECRET_KEY", ""),
        Port:         getEnvOrDefault("PORT", "8000"),
        AllowOrigins: getEnvOrDefault("ALLOWED_ORIGINS", "*"),
    }
}

func getEnvOrDefault(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}