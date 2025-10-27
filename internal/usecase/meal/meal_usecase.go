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

// GetAllMealItems 모든 급식 아이템 조회
func (u *MealUsecase) GetAllMealItems(ctx context.Context) ([]entities.MealItem, error) {
	return u.mealRepo.GetAll(ctx)
}

// GetMealItemByID ID로 급식 아이템 조회
func (u *MealUsecase) GetMealItemByID(ctx context.Context, id uint) (*entities.MealItem, error) {
	if id == 0 {
		return nil, errors.New("급식 아이템 ID는 필수입니다")
	}

	return u.mealRepo.GetByID(ctx, id)
}

// CreateMealItem 새로운 급식 아이템 생성
func (u *MealUsecase) CreateMealItem(ctx context.Context, req dto.MealItemCreateRequest) (*entities.MealItem, error) {
	// 급식 아이템 생성
	mealItem := &entities.MealItem{
		PetUserRoleID: req.PetUserRoleID,
		DataType:      req.DataType,
		MealType:      req.MealType,
		MealCategory:  req.MealCategory,
		Name:          req.Name,
		UnitType:      req.UnitType,
		IsDeleted:     false,
	}

	if err := u.mealRepo.Create(ctx, mealItem); err != nil {
		return nil, err
	}

	return mealItem, nil
}

// UpdateMealItem 급식 아이템 수정
func (u *MealUsecase) UpdateMealItem(ctx context.Context, id uint, req dto.MealItemUpdateRequest) error {
	if id == 0 {
		return errors.New("급식 아이템 ID는 필수입니다")
	}

	// 기존 급식 아이템 조회
	mealItem, err := u.mealRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// 수정할 필드만 업데이트
	if req.DataType != nil {
		mealItem.DataType = *req.DataType
	}
	if req.MealType != nil {
		mealItem.MealType = *req.MealType
	}
	if req.MealCategory != nil {
		mealItem.MealCategory = *req.MealCategory
	}
	if req.Name != nil {
		mealItem.Name = *req.Name
	}
	if req.UnitType != nil {
		mealItem.UnitType = *req.UnitType
	}

	return u.mealRepo.Update(ctx, id, mealItem)
}

// DeleteMealItem 급식 아이템 삭제
func (u *MealUsecase) DeleteMealItem(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("급식 아이템 ID는 필수입니다")
	}

	return u.mealRepo.Delete(ctx, id)
}

// GetMealItemsByType 타입별 급식 아이템 조회
func (u *MealUsecase) GetMealItemsByType(ctx context.Context, mealType entities.MealType) ([]entities.MealItem, error) {
	return u.mealRepo.GetByType(ctx, mealType)
}

// GetMealTypes 급식 타입 목록 조회
func (u *MealUsecase) GetMealTypes(ctx context.Context) []entities.MealType {
	return []entities.MealType{
		entities.MealTypeFeed,
		entities.MealTypeSnack,
	}
}
