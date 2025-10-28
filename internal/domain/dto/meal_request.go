package dto

import (
	"pet-manage-be/internal/domain/entities"
	"time"
)

// MealsCreateRequest 사료/간식 아이템 생성 요청
type MealsCreateRequest struct {
	PetUserRoleID uint              `json:"pet_user_role_id" validate:"required"`
	DataType      entities.DataType `json:"data_type" validate:"required,oneof=FIXED VARIATION"`
	MealType      entities.MealType `json:"meal_type" validate:"required"`
	MealCategory  string            `json:"meal_category" validate:"required"`
	Name          string            `json:"name" validate:"required"`
	UnitType      entities.UnitType `json:"unit_type" validate:"required"`
}

// MealsUpdateRequest 사료/간식 아이템 수정 요청
type MealsUpdateRequest struct {
	DataType     *entities.DataType `json:"data_type,omitempty"`
	MealType     *entities.MealType `json:"meal_type,omitempty"`
	MealCategory *string            `json:"meal_category,omitempty"`
	Name         *string            `json:"name,omitempty"`
	UnitType     *entities.UnitType `json:"unit_type,omitempty"`
}

// MealsUnitCreateRequest 사료/간식 단위 생성 요청
type MealsUnitCreateRequest struct {
	MealsID      uint                  `json:"meal_item_id" validate:"required"`
	MealUnitType entities.MealUnitType `json:"meal_item_type" validate:"required"`
	Unit         string                `json:"unit" validate:"required"`
	UnitValue    string                `json:"unit_value" validate:"required"`
}

// MealsUnitUpdateRequest 사료/간식 단위 수정 요청
type MealsUnitUpdateRequest struct {
	MealUnitType *entities.MealUnitType `json:"meal_item_type,omitempty"`
	Unit         *string                `json:"unit,omitempty"`
	UnitValue    *string                `json:"unit_value,omitempty"`
}

// MealHistoryCreateRequest 급여 기록 생성 요청
type MealHistoryCreateRequest struct {
	MealsID      uint              `json:"meal_item_id" validate:"required"`
	HistoryDate  time.Time         `json:"history_date" validate:"required"`
	MealType     entities.MealType `json:"meal_type" validate:"required"`
	Name         string            `json:"name" validate:"required"`
	Count        int               `json:"count" validate:"min=1"`
	MealCategory string            `json:"meal_category" validate:"required"`
	FeedAt       *time.Time        `json:"feed_at,omitempty"`
}

// MealHistoryUpdateRequest 급여 기록 수정 요청
type MealHistoryUpdateRequest struct {
	HistoryDate  *time.Time         `json:"history_date,omitempty"`
	MealType     *entities.MealType `json:"meal_type,omitempty"`
	Name         *string            `json:"name,omitempty"`
	Count        *int               `json:"count,omitempty"`
	MealCategory *string            `json:"meal_category,omitempty"`
	FeedAt       *time.Time         `json:"feed_at,omitempty"`
}

// MealHistoryUnitCreateRequest 급여 기록 단위 생성 요청
type MealHistoryUnitCreateRequest struct {
	MealHistoryID uint                  `json:"meal_histories_id" validate:"required"`
	MealUnitType  entities.MealUnitType `json:"meal_item_type" validate:"required"`
	Unit          string                `json:"unit" validate:"required"`
	UnitValue     string                `json:"unit_value" validate:"required"`
}

// MealHistoryUnitUpdateRequest 급여 기록 단위 수정 요청
type MealHistoryUnitUpdateRequest struct {
	MealUnitType *entities.MealUnitType `json:"meal_item_type,omitempty"`
	Unit         *string                `json:"unit,omitempty"`
	UnitValue    *string                `json:"unit_value,omitempty"`
}
