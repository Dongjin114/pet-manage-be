package dto

import (
	"pet-manage-be/internal/domain/entities"
	"time"
)

// MealsResponse 사료/간식 아이템 응답
type MealsResponse struct {
	ID            uint              `json:"id"`
	PetUserRoleID uint              `json:"pet_user_role_id"`
	DataType      entities.DataType `json:"data_type"`
	MealType      entities.MealType `json:"meal_type"`
	MealCategory  string            `json:"meal_category"`
	Name          string            `json:"name"`
	UnitType      entities.UnitType `json:"unit_type"`
	CreatedAt     time.Time         `json:"created_at"`
	ModifiedAt    *time.Time        `json:"modified_at"`
	DeletedAt     *time.Time        `json:"deleted_at"`
	IsDeleted     bool              `json:"is_deleted"`
}

// MealsWithUnitsResponse 사료/간식 아이템과 단위 정보를 포함한 응답
type MealsWithUnitsResponse struct {
	MealsResponse
	Units []MealsUnitResponse `json:"units"`
}

// MealsUnitResponse 사료/간식 단위 응답
type MealsUnitResponse struct {
	ID           uint                  `json:"id"`
	MealsID      uint                  `json:"meal_item_id"`
	MealUnitType entities.MealUnitType `json:"meal_item_type"`
	Unit         string                `json:"unit"`
	UnitValue    string                `json:"unit_value"`
}

// MealHistoryResponse 급여 기록 응답
type MealHistoryResponse struct {
	ID           uint              `json:"id"`
	MealsID      uint              `json:"meal_item_id"`
	HistoryDate  time.Time         `json:"history_date"`
	MealType     entities.MealType `json:"meal_type"`
	Name         string            `json:"name"`
	Count        int               `json:"count"`
	MealCategory string            `json:"meal_category"`
	FeedAt       *time.Time        `json:"feed_at"`
	ModifiedAt   *time.Time        `json:"modified_at"`
	DeletedAt    *time.Time        `json:"deleted_at"`
	IsDeleted    bool              `json:"is_deleted"`
}

// MealHistoryWithUnitsResponse 급여 기록과 단위 정보를 포함한 응답
type MealHistoryWithUnitsResponse struct {
	MealHistoryResponse
	Units []MealHistoryUnitResponse `json:"units"`
}

// MealHistoryUnitResponse 급여 기록 단위 응답
type MealHistoryUnitResponse struct {
	ID            uint                  `json:"id"`
	MealHistoryID uint                  `json:"meal_histories_id"`
	MealUnitType  entities.MealUnitType `json:"meal_item_type"`
	Unit          string                `json:"unit"`
	UnitValue     string                `json:"unit_value"`
}

// MealTypesResponse 급식 타입 목록 응답
type MealTypesResponse struct {
	Types []entities.MealType `json:"types"`
	Count int                 `json:"count"`
}

// DataTypesResponse 데이터 타입 목록 응답
type DataTypesResponse struct {
	Types []entities.DataType `json:"types"`
	Count int                 `json:"count"`
}

// UnitTypesResponse 단위 타입 목록 응답
type UnitTypesResponse struct {
	Types []entities.UnitType `json:"types"`
	Count int                 `json:"count"`
}

// MealUnitTypesResponse 식이 타입 목록 응답
type MealUnitTypesResponse struct {
	Types []entities.MealUnitType `json:"types"`
	Count int                     `json:"count"`
}
