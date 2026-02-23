package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port          string
	DatabaseURL   string
	UploadPath    string
	MaxUploadSize int64
	BatchSize     int
	MaxGoroutines int
	UseS3         bool
	S3Bucket      string
	AWSRegion     string
	LogLevel      string
}

func Load() *Config {
	return &Config{
		Port:          getEnv("PORT", "8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "pricing.db"),
		UploadPath:    getEnv("UPLOAD_PATH", "./uploads"),
		MaxUploadSize: getEnvAsInt64("MAX_UPLOAD_SIZE", 10485760), // 10MB
		BatchSize:     getEnvAsInt("BATCH_SIZE", 1000),
		MaxGoroutines: getEnvAsInt("MAX_GOROUTINES", 10),
		UseS3:         getEnvAsBool("USE_S3", false),
		S3Bucket:      getEnv("S3_BUCKET", ""),
		AWSRegion:     getEnv("AWS_REGION", "us-east-1"),
		LogLevel:      getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsInt64(key string, defaultVal int64) int64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsBool(key string, defaultVal bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func (c *Config) Validate() {
	if c.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	// Create upload directory if it doesn't exist
	if err := os.MkdirAll(c.UploadPath, os.ModePerm); err != nil {
		log.Fatalf("Failed to create upload directory: %v", err)
	}

	log.Printf("Configuration loaded successfully")
	log.Printf("Port: %s", c.Port)
	log.Printf("Upload Path: %s", c.UploadPath)
	log.Printf("Batch Size: %d", c.BatchSize)
}
