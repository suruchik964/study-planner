package config

import "os"

// Config holds all environment-based settings in one place.
type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	JWTSecret  string
}

func Load() Config {
	return Config{
		AppPort:    getEnvWithFallback([]string{"APP_PORT", "PORT"}, "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "planner_user"),
		DBPassword: getEnv("DB_PASSWORD", "planner_password"),
		DBName:     getEnv("DB_NAME", "study_planner"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		JWTSecret:  getEnv("JWT_SECRET", "super-secret-study-planner-key"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getEnvWithFallback(keys []string, fallback string) string {
	for _, key := range keys {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}
	return fallback
}
