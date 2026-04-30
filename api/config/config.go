package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	Port                     string
	Host                     string
	DB_url                   string
	StorageURL               string
	StorageAccessKeyID       string
	StorageSecretAccessKey   string
	StorageRegion            string
	StorageProofToken        string
	StartTime                time.Time
	RedisAddr                string
	RedisPassword            string
	RedisDB                  int
	AllowedRequestsPerMinute int
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvAsInt(name string, defaultVal int) int {
	if valueStr, exists := os.LookupEnv(name); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultVal
}

func Init() {
	var err error
	Port = GetEnv("PORT", "8081")
	Host = GetEnv("HOST", "0.0.0.0")

	user := GetEnv("DB_USER", "workerbee")
	password := GetEnv("DB_PASSWORD", "")
	port := GetEnv("DB_PORT", "5432")
	db_name := GetEnv("DB", "workerbee")
	db_host := GetEnv("DB_HOST", "localhost")
	DB_url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, db_host, port, db_name)
	StorageURL = GetEnv("S3_URL", GetEnv("DO_URL", ""))
	StorageAccessKeyID = GetEnv("S3_ACCESS_KEY_ID", GetEnv("DO_ACCESS_KEY_ID", ""))
	StorageSecretAccessKey = GetEnv("S3_SECRET_ACCESS_KEY", GetEnv("DO_SECRET_ACCESS_KEY", ""))
	StorageRegion = GetEnv("S3_REGION", "us-east-1")
	StorageProofToken = GetEnv("STORAGE_PROOF_TOKEN", "")
	RedisAddr = GetEnv("REDIS_ADDR", "localhost:6379")
	RedisPassword = GetEnv("REDIS_PASSWORD", "")
	RedisDB = GetEnvAsInt("REDIS_DB", 0)

	StartTime = time.Now()
	RateLimitRoofStr := GetEnv("ALLOWED_PROTECTED_REQUESTS", "25")
	_, err = fmt.Sscanf(RateLimitRoofStr, "%d", &AllowedRequestsPerMinute)
	if err != nil || AllowedRequestsPerMinute <= 0 {
		AllowedRequestsPerMinute = 25
	}
}
