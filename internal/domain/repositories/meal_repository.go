package repositories

import (
	"context"
	"pet-manage-be/internal/domain/entities"
)

// MealRepository 급식 리포지토리 인터페이스
type MealRepository interface {
	// GetAll 모든 급식 아이템 조회
	GetAll(ctx context.Context) ([]entities.MealItem, error)

	// GetByID ID로 급식 아이템 조회
	GetByID(ctx context.Context, id uint) (*entities.MealItem, error)

	// Create 급식 아이템 생성
	Create(ctx context.Context, meal *entities.MealItem) error

	// Update 급식 아이템 수정
	Update(ctx context.Context, id uint, meal *entities.MealItem) error

	// Delete 급식 아이템 삭제
	Delete(ctx context.Context, id uint) error

	// GetByType 타입별 급식 아이템 조회
	GetByType(ctx context.Context, mealType entities.MealType) ([]entities.MealItem, error)
}
