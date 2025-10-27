package migrations

import (
	"database/sql"
	"fmt"
	"log"
)

// RunMigrations 마이그레이션 실행
func RunMigrations(db *sql.DB) error {
	log.Println("PostgreSQL 데이터베이스 마이그레이션 시작...")

	migrations := []struct {
		name string
		sql  string
	}{
		{
			name: "1",
			sql: `
				CREATE TABLE IF NOT EXISTS owners (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					name VARCHAR(100) NOT NULL,
					email VARCHAR(255) UNIQUE NOT NULL,
					phone VARCHAR(20),
					address TEXT,
					created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
					modified_at TIMESTAMPTZ,
					deleted_at TIMESTAMPTZ,
					is_deleted BOOLEAN DEFAULT FALSE
				);
			`,
		},
		{
			name: "2",
			sql: `
				CREATE TABLE IF NOT EXISTS pets (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					owner_id BIGINT NOT NULL REFERENCES owners(id),
					name VARCHAR(100) NOT NULL,
					species VARCHAR(50) NOT NULL,
					breed VARCHAR(100),
					age INTEGER,
					weight DECIMAL(5,2),
					gender VARCHAR(10),
					created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
					modified_at TIMESTAMPTZ,
					deleted_at TIMESTAMPTZ,
					is_deleted BOOLEAN DEFAULT FALSE
				);
			`,
		},
		{
			name: "3",
			sql: `
				CREATE TABLE IF NOT EXISTS meal_items (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					pet_user_role_id BIGINT NOT NULL,
					data_type VARCHAR(20) NOT NULL CHECK (data_type IN ('FIXED', 'VARIATION')),
					meal_type VARCHAR(20) NOT NULL CHECK (meal_type IN ('사료', '간식')),
					meal_category VARCHAR(100) NOT NULL,
					name VARCHAR(200) NOT NULL,
					unit_type VARCHAR(10) NOT NULL CHECK (unit_type IN ('g', 'ml', '개', '포')),
					created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
					modified_at TIMESTAMPTZ,
					deleted_at TIMESTAMPTZ,
					is_deleted BOOLEAN DEFAULT FALSE
				);
			`,
		},
		{
			name: "4",
			sql: `
				CREATE TABLE IF NOT EXISTS meal_item_units (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					meal_item_id BIGINT NOT NULL REFERENCES meal_items(id),
					unit_type VARCHAR(10) NOT NULL CHECK (unit_type IN ('g', 'ml', '개', '포')),
					amount DECIMAL(10,2) NOT NULL,
					created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
					modified_at TIMESTAMPTZ,
					deleted_at TIMESTAMPTZ,
					is_deleted BOOLEAN DEFAULT FALSE
				);
			`,
		},
		{
			name: "5",
			sql: `
				CREATE TABLE IF NOT EXISTS meal_histories (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					pet_id BIGINT NOT NULL REFERENCES pets(id),
					meal_item_id BIGINT NOT NULL REFERENCES meal_items(id),
					meal_date DATE NOT NULL,
					meal_time TIME NOT NULL,
					amount DECIMAL(10,2) NOT NULL,
					unit VARCHAR(10) NOT NULL,
					notes TEXT,
					created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
					modified_at TIMESTAMPTZ,
					deleted_at TIMESTAMPTZ,
					is_deleted BOOLEAN DEFAULT FALSE
				);
			`,
		},
		{
			name: "6",
			sql: `
				CREATE TABLE IF NOT EXISTS meal_history_units (
					id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
					meal_history_id BIGINT NOT NULL REFERENCES meal_histories(id),
					unit_type VARCHAR(10) NOT NULL CHECK (unit_type IN ('g', 'ml', '개', '포')),
					amount DECIMAL(10,2) NOT NULL,
					created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
					modified_at TIMESTAMPTZ,
					deleted_at TIMESTAMPTZ,
					is_deleted BOOLEAN DEFAULT FALSE
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