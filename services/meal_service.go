package services

import (
	"errors"
	"pet-manage-be/database"
	"pet-manage-be/models"

	"gorm.io/gorm"
)

// MealService handles meal-related business logic
type MealService struct {
	db *gorm.DB
}

// NewMealService creates a new meal service
func NewMealService() *MealService {
	return &MealService{
		db: database.DB,
	}
}

// CreateMealItem creates a new meal item with validation
func (s *MealService) CreateMealItem(mealItem *models.MealItem) error {
	// 1. MealType 유효성 검사
	if !mealItem.DataType.IsValid() {
		return errors.New("invalid meal type: must be FIXED or VARIATION")
	}

	// 2. 추가 비즈니스 로직 검증
	if mealItem.Name == "" {
		return errors.New("meal name is required")
	}

	if mealItem.Amount < 0 {
		return errors.New("meal amount cannot be negative")
	}

	// 3. 데이터베이스에 저장
	return s.db.Create(mealItem).Error
}

// UpdateMealItem updates an existing meal item
func (s *MealService) UpdateMealItem(id uint, updates map[string]interface{}) error {
	// DataType이 포함된 경우 유효성 검사
	if dataType, exists := updates["data_type"]; exists {
		if mealType, ok := dataType.(models.MealType); ok {
			if !mealType.IsValid() {
				return errors.New("invalid meal type: must be FIXED or VARIATION")
			}
		}
	}

	return s.db.Model(&models.MealItem{}).Where("id = ?", id).Updates(updates).Error
}

// GetMealItemsByType gets meal items by type
func (s *MealService) GetMealItemsByType(mealType models.MealType) ([]models.MealItem, error) {
	// 유효성 검사
	if !mealType.IsValid() {
		return nil, errors.New("invalid meal type")
	}

	var mealItems []models.MealItem
	err := s.db.Where("data_type = ?", mealType).Find(&mealItems).Error
	return mealItems, err
}

// GetAllMealItems gets all meal items
func (s *MealService) GetAllMealItems() ([]models.MealItem, error) {
	var mealItems []models.MealItem
	err := s.db.Find(&mealItems).Error
	return mealItems, err
}

// GetMealItemByID gets a meal item by ID
func (s *MealService) GetMealItemByID(id uint) (*models.MealItem, error) {
	var mealItem models.MealItem
	err := s.db.First(&mealItem, id).Error
	if err != nil {
		return nil, err
	}
	return &mealItem, nil
}

// DeleteMealItem deletes a meal item
func (s *MealService) DeleteMealItem(id uint) error {
	return s.db.Delete(&models.MealItem{}, id).Error
}
