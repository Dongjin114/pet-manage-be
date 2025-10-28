package repository

import (
	"context"
	"database/sql"
	"fmt"
	"pet-manage-be/internal/domain/entities"
)

// MealRepository PostgreSQL 구현
type MealRepository struct {
	db *sql.DB
}

// NewMealRepository 새로운 MealRepository 생성
func NewMealRepository() *MealRepository {
	return &MealRepository{}
}

// SetDB 데이터베이스 설정
func (r *MealRepository) SetDB(db *sql.DB) {
	r.db = db
}

// GetAll 모든 급식 아이템 조회
func (r *MealRepository) GetAll(ctx context.Context) ([]entities.Meals, error) {
	query := `
		SELECT id, pet_user_role_id, data_type, meal_type, meal_category, name, unit_type, 
		       created_at, modified_at, deleted_at, is_deleted
		FROM meal_items 
		WHERE is_deleted = FALSE
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("급식 아이템 조회 실패: %w", err)
	}
	defer rows.Close()

	var items []entities.Meals
	for rows.Next() {
		var item entities.Meals
		var modifiedAt, deletedAt sql.NullTime

		err := rows.Scan(
			&item.ID, &item.PetUserRoleID, &item.DataType, &item.MealType,
			&item.MealCategory, &item.Name, &item.UnitType,
			&item.CreatedAt, &modifiedAt, &deletedAt, &item.IsDeleted,
		)
		if err != nil {
			return nil, fmt.Errorf("급식 아이템 스캔 실패: %w", err)
		}

		if modifiedAt.Valid {
			item.ModifiedAt = &modifiedAt.Time
		}
		if deletedAt.Valid {
			item.DeletedAt = &deletedAt.Time
		}

		items = append(items, item)
	}

	return items, nil
}

// GetByID ID로 급식 아이템 조회
func (r *MealRepository) GetByID(ctx context.Context, id uint) (*entities.Meals, error) {
	query := `
		SELECT id, pet_user_role_id, data_type, meal_type, meal_category, name, unit_type,
		       created_at, modified_at, deleted_at, is_deleted
		FROM meal_items 
		WHERE id = $1 AND is_deleted = FALSE
	`

	var item entities.Meals
	var modifiedAt, deletedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&item.ID, &item.PetUserRoleID, &item.DataType, &item.MealType,
		&item.MealCategory, &item.Name, &item.UnitType,
		&item.CreatedAt, &modifiedAt, &deletedAt, &item.IsDeleted,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("급식 아이템을 찾을 수 없습니다")
		}
		return nil, fmt.Errorf("급식 아이템 조회 실패: %w", err)
	}

	if modifiedAt.Valid {
		item.ModifiedAt = &modifiedAt.Time
	}
	if deletedAt.Valid {
		item.DeletedAt = &deletedAt.Time
	}

	return &item, nil
}

// Create 급식 아이템 생성
func (r *MealRepository) Create(ctx context.Context, item *entities.Meals) error {
	query := `
		INSERT INTO meal_items (pet_user_role_id, data_type, meal_type, meal_category, name, unit_type)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	err := r.db.QueryRowContext(ctx, query,
		item.PetUserRoleID, item.DataType, item.MealType,
		item.MealCategory, item.Name, item.UnitType,
	).Scan(&item.ID, &item.CreatedAt)

	if err != nil {
		return fmt.Errorf("급식 아이템 생성 실패: %w", err)
	}

	return nil
}

// Update 급식 아이템 수정
func (r *MealRepository) Update(ctx context.Context, id uint, item *entities.Meals) error {
	query := `
		UPDATE meal_items 
		SET pet_user_role_id = $2, data_type = $3, meal_type = $4, 
		    meal_category = $5, name = $6, unit_type = $7, modified_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND is_deleted = FALSE
	`

	result, err := r.db.ExecContext(ctx, query,
		id, item.PetUserRoleID, item.DataType, item.MealType,
		item.MealCategory, item.Name, item.UnitType,
	)

	if err != nil {
		return fmt.Errorf("급식 아이템 수정 실패: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("급식 아이템 수정 결과 확인 실패: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("급식 아이템을 찾을 수 없습니다")
	}

	return nil
}

// Delete 급식 아이템 삭제 (소프트 삭제)
func (r *MealRepository) Delete(ctx context.Context, id uint) error {
	query := `
		UPDATE meal_items 
		SET is_deleted = TRUE, deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND is_deleted = FALSE
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("급식 아이템 삭제 실패: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("급식 아이템 삭제 결과 확인 실패: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("급식 아이템을 찾을 수 없습니다")
	}

	return nil
}

// GetByType 타입별 급식 아이템 조회
func (r *MealRepository) GetByType(ctx context.Context, mealType entities.MealType) ([]entities.Meals, error) {
	query := `
		SELECT id, pet_user_role_id, data_type, meal_type, meal_category, name, unit_type,
		       created_at, modified_at, deleted_at, is_deleted
		FROM meal_items 
		WHERE meal_type = $1 AND is_deleted = FALSE
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, mealType)
	if err != nil {
		return nil, fmt.Errorf("타입별 급식 아이템 조회 실패: %w", err)
	}
	defer rows.Close()

	var items []entities.Meals
	for rows.Next() {
		var item entities.Meals
		var modifiedAt, deletedAt sql.NullTime

		err := rows.Scan(
			&item.ID, &item.PetUserRoleID, &item.DataType, &item.MealType,
			&item.MealCategory, &item.Name, &item.UnitType,
			&item.CreatedAt, &modifiedAt, &deletedAt, &item.IsDeleted,
		)
		if err != nil {
			return nil, fmt.Errorf("타입별 급식 아이템 스캔 실패: %w", err)
		}

		if modifiedAt.Valid {
			item.ModifiedAt = &modifiedAt.Time
		}
		if deletedAt.Valid {
			item.DeletedAt = &deletedAt.Time
		}

		items = append(items, item)
	}

	return items, nil
}
