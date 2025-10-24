package repository

import (
	"context"
	"pet-manage-be/internal/domain/entities"
	"pet-manage-be/internal/domain/repositories"
)

// mealRepositoryImpl 급식 리포지토리 구현체
type mealRepositoryImpl struct {
	// TODO: 데이터베이스 연결 추가
}

// NewMealRepository 새로운 급식 리포지토리 생성
func NewMealRepository() repositories.MealRepository {
	return &mealRepositoryImpl{}
}

// GetAll 모든 급식 아이템 조회
func (r *mealRepositoryImpl) GetAll(ctx context.Context) ([]entities.MealItem, error) {
	// TODO: 데이터베이스에서 조회
	// 현재는 목 데이터 반환
	mealItems := []entities.MealItem{
		{
			ID:       1,
			DataType: entities.MealTypeFixed,
			Name:     "Royal Canin",
			Amount:   100.0,
			Unit:     "g",
		},
		{
			ID:       2,
			DataType: entities.MealTypeVariation,
			Name:     "Chicken Treats",
			Amount:   50.0,
			Unit:     "g",
		},
	}
	return mealItems, nil
}

// GetByID ID로 급식 아이템 조회
func (r *mealRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.MealItem, error) {
	// TODO: 데이터베이스에서 조회
	// 현재는 목 데이터 반환
	mealItem := &entities.MealItem{
		ID:       id,
		DataType: entities.MealTypeFixed,
		Name:     "Royal Canin",
		Amount:   100.0,
		Unit:     "g",
	}
	return mealItem, nil
}

// Create 급식 아이템 생성
func (r *mealRepositoryImpl) Create(ctx context.Context, meal *entities.MealItem) error {
	// TODO: 데이터베이스에 저장
	// 현재는 성공으로 처리
	return nil
}

// Update 급식 아이템 수정
func (r *mealRepositoryImpl) Update(ctx context.Context, id uint, meal *entities.MealItem) error {
	// TODO: 데이터베이스에서 수정
	// 현재는 성공으로 처리
	return nil
}

// Delete 급식 아이템 삭제
func (r *mealRepositoryImpl) Delete(ctx context.Context, id uint) error {
	// TODO: 데이터베이스에서 삭제
	// 현재는 성공으로 처리
	return nil
}

// GetByType 타입별 급식 아이템 조회
func (r *mealRepositoryImpl) GetByType(ctx context.Context, mealType entities.MealType) ([]entities.MealItem, error) {
	// TODO: 데이터베이스에서 조회
	// 현재는 목 데이터 반환
	mealItems := []entities.MealItem{
		{
			ID:       1,
			DataType: mealType,
			Name:     "Royal Canin",
			Amount:   100.0,
			Unit:     "g",
		},
	}
	return mealItems, nil
}
