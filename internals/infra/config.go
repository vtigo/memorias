package infra

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	BaseURL  string
	Port     int
	LogLevel slog.Level
	S3Bucket *string
}

func LoadConfig() (*Config, error) {
	port, err := strconv.Atoi(getEnv("PORT", "3333"))
	if err != nil {
		return nil, fmt.Errorf("invalid PORT: %w", err)
	}

	logLevel, err := parseLogLevel(getEnv("LOG_LEVEL", "info"))
	if err != nil {
		return nil, fmt.Errorf("invalid LOG_LEVEL: %w", err)
	}
	
	config := &Config{
		BaseURL:  getEnv("BASE_URL", "0.0.0.0"),
		Port:     port,
		LogLevel: logLevel,
	}

	if bucket, ok := os.LookupEnv("S3_BUCKET"); ok {
		config.S3Bucket = &bucket
	}

	return config, nil
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func parseLogLevel(v string) (slog.Level, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(v)); err != nil {
		return level, fmt.Errorf("unknown log level %q", v)
	}
	return level, nil
}

