package migrations

import (
	"database/sql"
	"fmt"
	"log"
)

// 공통 컬럼 정의
const (
	CommonColumns = `
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		created_by BIGINT,
		modified_at TIMESTAMPTZ,
		modified_by BIGINT,
		deleted_at TIMESTAMPTZ,
		deleted_by BIGINT,
		is_deleted BOOLEAN DEFAULT FALSE
	`
)

// RunMigrations 마이그레이션 실행
func RunMigrations(db *sql.DB) error {
	log.Println("PostgreSQL 데이터베이스 마이그레이션 시작...")

	migrations := []struct {
		name string
		sql  string
	}{
		{
			name: "사용자 테이블",
			sql: `
				CREATE TABLE IF NOT EXISTS owners (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					email VARCHAR(100) NOT NULL,
					name VARCHAR(255) UNIQUE NOT NULL,
					` + CommonColumns + `
				);
			`,
		},
		{
			name: "2",
			sql: `
				CREATE TABLE IF NOT EXISTS pets (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					owner_id BIGINT NOT NULL,
					name VARCHAR(100) NOT NULL,
					species VARCHAR(50) NOT NULL,
					breed VARCHAR(100),
					age INTEGER,
					weight DECIMAL(5,2),
					gender VARCHAR(10),
					` + CommonColumns + `
				);
			`,
		},
		{
			name: "3",
			sql: `
				CREATE TABLE IF NOT EXISTS meals (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					pet_user_role_id BIGINT NOT NULL,
					data_type VARCHAR(20) NOT NULL CHECK (data_type IN ('FIXED', 'VARIATION')),
					meal_type VARCHAR(20) NOT NULL CHECK (meal_type IN ('사료', '간식')),
					meal_category VARCHAR(100) NOT NULL,
					name VARCHAR(200) NOT NULL,
					unit_type VARCHAR(10) NOT NULL CHECK (unit_type IN ('g', 'ml', '개', '포')),
					` + CommonColumns + `
				);
			`,
		},
		{
			name: "4",
			sql: `
				CREATE TABLE IF NOT EXISTS meal_units (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					meal_id BIGINT NOT NULL,
					unit_type VARCHAR(10) NOT NULL CHECK (unit_type IN ('g', 'ml', '개', '포')),
					amount DECIMAL(10,2) NOT NULL,
					` + CommonColumns + `
				);
			`,
		},
		{
			name: "5",
			sql: `
				CREATE TABLE IF NOT EXISTS meal_histories (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					pet_id BIGINT NOT NULL,
					meal_id BIGINT NOT NULL,
					meal_date DATE NOT NULL,
					meal_time TIME NOT NULL,
					amount DECIMAL(10,2) NOT NULL,
					unit VARCHAR(10) NOT NULL,
					notes TEXT,
					` + CommonColumns + `
				);
			`,
		},
		{
			name: "6",
			sql: `
				CREATE TABLE IF NOT EXISTS meal_history_units (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					meal_history_id BIGINT NOT NULL,
					unit_type VARCHAR(10) NOT NULL CHECK (unit_type IN ('g', 'ml', '개', '포')),
					amount DECIMAL(10,2) NOT NULL,
					` + CommonColumns + `
				);
			`,
		},
	}

	for _, migration := range migrations {
		log.Printf("마이그레이션 %s 실행 중...", migration.name)
		if _, err := db.Exec(migration.sql); err != nil {
			return fmt.Errorf("마이그레이션 %s 실행 실패: %w", migration.name, err)
		}
	}

	log.Println("모든 마이그레이션 완료")
	return nil
}
