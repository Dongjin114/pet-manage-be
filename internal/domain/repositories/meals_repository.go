package repositories

import (
	"context"
	"pet-manage-be/internal/domain/entities"
)

// MealRepository 급식 리포지토리 인터페이스
type MealRepository interface {
	// GetAll 모든 급식 아이템 조회
	GetAll(ctx context.Context) ([]entities.Meals, error)

	// GetByID ID로 급식 아이템 조회
	GetByID(ctx context.Context, id uint) (*entities.Meals, error)

	// Create 급식 아이템 생성
	Create(ctx context.Context, meal *entities.Meals) error

	// Update 급식 아이템 수정
	Update(ctx context.Context, id uint, meal *entities.Meals) error

	// Delete 급식 아이템 삭제
	Delete(ctx context.Context, id uint) error

	// GetByType 타입별 급식 아이템 조회
	GetByType(ctx context.Context, mealType entities.MealType) ([]entities.Meals, error)
}
