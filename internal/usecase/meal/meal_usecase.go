package meal

import (
	"context"
	"errors"
	"pet-manage-be/internal/domain/dto"
	"pet-manage-be/internal/domain/entities"
	"pet-manage-be/internal/domain/repositories"
)

// MealUsecase 급식 유스케이스
type MealUsecase struct {
	mealRepo repositories.MealRepository
}

// NewMealUsecase 새로운 급식 유스케이스 생성
func NewMealUsecase(mealRepo repositories.MealRepository) *MealUsecase {
	return &MealUsecase{
		mealRepo: mealRepo,
	}
}

// GetAllMealss 모든 급식 아이템 조회
func (u *MealUsecase) GetAllMealss(ctx context.Context) ([]entities.Meals, error) {
	return u.mealRepo.GetAll(ctx)
}

// GetMealsByID ID로 급식 아이템 조회
func (u *MealUsecase) GetMealsByID(ctx context.Context, id uint) (*entities.Meals, error) {
	if id == 0 {
		return nil, errors.New("급식 아이템 ID는 필수입니다")
	}

	return u.mealRepo.GetByID(ctx, id)
}

// CreateMeals 새로운 급식 아이템 생성
func (u *MealUsecase) CreateMeals(ctx context.Context, req dto.MealsCreateRequest) (*entities.Meals, error) {
	// 급식 아이템 생성
	Meals := &entities.Meals{
		PetUserRoleID: req.PetUserRoleID,
		DataType:      req.DataType,
		MealType:      req.MealType,
		MealCategory:  req.MealCategory,
		Name:          req.Name,
		UnitType:      req.UnitType,
		IsDeleted:     false,
	}

	if err := u.mealRepo.Create(ctx, Meals); err != nil {
		return nil, err
	}

	return Meals, nil
}

// UpdateMeals 급식 아이템 수정
func (u *MealUsecase) UpdateMeals(ctx context.Context, id uint, req dto.MealsUpdateRequest) error {
	if id == 0 {
		return errors.New("급식 아이템 ID는 필수입니다")
	}

	// 기존 급식 아이템 조회
	Meals, err := u.mealRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// 수정할 필드만 업데이트
	if req.DataType != nil {
		Meals.DataType = *req.DataType
	}
	if req.MealType != nil {
		Meals.MealType = *req.MealType
	}
	if req.MealCategory != nil {
		Meals.MealCategory = *req.MealCategory
	}
	if req.Name != nil {
		Meals.Name = *req.Name
	}
	if req.UnitType != nil {
		Meals.UnitType = *req.UnitType
	}

	return u.mealRepo.Update(ctx, id, Meals)
}

// DeleteMeals 급식 아이템 삭제
func (u *MealUsecase) DeleteMeals(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("급식 아이템 ID는 필수입니다")
	}

	return u.mealRepo.Delete(ctx, id)
}

// GetMealssByType 타입별 급식 아이템 조회
func (u *MealUsecase) GetMealssByType(ctx context.Context, mealType entities.MealType) ([]entities.Meals, error) {
	return u.mealRepo.GetByType(ctx, mealType)
}

// GetMealTypes 급식 타입 목록 조회
func (u *MealUsecase) GetMealTypes(ctx context.Context) []entities.MealType {
	return []entities.MealType{
		entities.MealTypeFeed,
		entities.MealTypeSnack,
	}
}
