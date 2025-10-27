package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config 애플리케이션 설정
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig 서버 설정
type ServerConfig struct {
	Port string
	Host string
}

// DatabaseConfig 데이터베이스 설정
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Load 설정 로드
func Load() *Config {
	// .env 파일 로드
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env 파일을 찾을 수 없습니다:", err)
	}

	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Host: getEnv("HOST", "0.0.0.0"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "user"),
			Password: getEnv("DB_PASSWORD", "12345678"),
			Name:     getEnv("DB_NAME", "petdb"),
		},
	}
}

// getEnv 환경변수 가져오기
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
