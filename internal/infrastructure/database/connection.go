package database

import (
	"database/sql"
	"fmt"
	"log"
	"pet-manage-be/internal/infrastructure/config"

	_ "github.com/lib/pq"
)

// DB 데이터베이스 연결
var DB *sql.DB

// Connect 데이터베이스 연결 (PostgreSQL)
func Connect(cfg *config.DatabaseConfig) error {
	// PostgreSQL 연결 문자열 생성
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	// 데이터베이스 연결
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("PostgreSQL 데이터베이스 연결 실패: %w", err)
	}

	// 연결 테스트
	if err := db.Ping(); err != nil {
		return fmt.Errorf("PostgreSQL 데이터베이스 연결 테스트 실패: %w", err)
	}

	// 연결 풀 설정
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	DB = db
	log.Printf("PostgreSQL 데이터베이스 연결 성공: %s:%s/%s", cfg.Host, cfg.Port, cfg.Name)
	return nil
}

// Close 데이터베이스 연결 종료
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// GetDB 데이터베이스 인스턴스 반환
func GetDB() *sql.DB {
	return DB
}
