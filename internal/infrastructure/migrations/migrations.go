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
				CREATE TABLE IF NOT EXISTS USERS (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					email VARCHAR(100) NOT NULL,
					name VARCHAR(255) UNIQUE NOT NULL,
					` + CommonColumns + `
				);
			`,
		},
		{
			name: "반려동물 테이블",
			sql: `
				CREATE TABLE IF NOT EXISTS PETS (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					registration_type VARCHAR(50) NOT NULL CHECK (registration_type IN ('OFFICIAL', 'MANUAL')),
					reg_number VARCHAR(100) NOT NULL,
					owner_user_id BIGINT NOT NULL,
					name VARCHAR(100) NOT NULL,
					animal_type VARCHAR(50) NOT NULL,
					species VARCHAR(50) NOT NULL,
					age INTEGER,
					gender VARCHAR(10),
					` + CommonColumns + `
				);
			`,
		},

		{
			name: "반려동물 유저 권한 테이블",
			sql: `
				CREATE TABLE IF NOT EXISTS PET_USER_ROLES (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					pet_id BIGINT NOT NULL,
					user_id BIGINT NOT NULL,
					role VARCHAR(50) NOT NULL CHECK (role IN ('OWNER', 'FAMILY_MEMBER', 'GUEST', 'CARETAKER')),
					` + CommonColumns + `
				);
			`,
		},
		{
			name: "식이 테이블",
			sql: `
				CREATE TABLE IF NOT EXISTS MEALS (
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
			name: "식이 단위 테이블",
			sql: `
				CREATE TABLE IF NOT EXISTS meal_units (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					meal_id BIGINT NOT NULL,
					meal_unit_type VARCHAR(50) NOT NULL,
					unit_type VARCHAR(10) NOT NULL CHECK (unit_type IN ('g', 'ml', '개', '포')),
					amount DECIMAL(10,2) NOT NULL,
					` + CommonColumns + `
				);
			`,
		},
		{
			name: "식이 기록 데이터 테이블",
			sql: `
				CREATE TABLE IF NOT EXISTS meal_histories (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					pet_id BIGINT NOT NULL,
					meal_id BIGINT NOT NULL,
					feed_date DATE NOT NULL,
					meal_type TIME NOT NULL,
					amount DECIMAL(10,2) NOT NULL,
					meal_category VARCHAR(100) NOT NULL,
					notes TEXT,
					` + CommonColumns + `
				);
			`,
		},
		{
			name: "식이 기록 단위 테이블",
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
