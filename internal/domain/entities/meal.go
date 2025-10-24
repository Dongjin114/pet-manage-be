package entities

import (
	"time"
)

// DataType 데이터 타입
type DataType string

const (
	DataTypeFixed     DataType = "FIXED"
	DataTypeVariation DataType = "VARIATION"
)

// MealType 급식 타입 (사료/간식 구분)
type MealType string

const (
	MealTypeFeed  MealType = "사료"
	MealTypeSnack MealType = "간식"
)

// UnitType 단위 타입
type UnitType string

const (
	UnitTypeGram     UnitType = "g"
	UnitTypeKilogram UnitType = "kg"
	UnitTypePiece    UnitType = "개"
	UnitTypeCup      UnitType = "컵"
)

// MealItemType 식이 타입
type MealItemType string

const (
	MealItemTypeRawFeeding MealItemType = "원체급여량"
	MealItemTypeCalorie    MealItemType = "칼로리"
	MealItemTypeProtein    MealItemType = "단백질"
	MealItemTypeFat        MealItemType = "지방"
	MealItemTypeCarb       MealItemType = "탄수화물"
)

// MealItem 사료/간식 아이템 엔티티
type MealItem struct {
	ID            uint       `json:"id" db:"id"`
	PetUserRoleID uint       `json:"pet_user_role_id" db:"pet_user_role_id"`
	DataType      DataType   `json:"data_type" db:"data_type"`
	MealType      MealType   `json:"meal_type" db:"meal_type"`
	MealCategory  string     `json:"meal_category" db:"meal_category"`
	Name          string     `json:"name" db:"name"`
	UnitType      UnitType   `json:"unit_type" db:"unit_type"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	ModifiedAt    *time.Time `json:"modified_at" db:"modified_at"`
	DeletedAt     *time.Time `json:"deleted_at" db:"deleted_at"`
	IsDeleted     bool       `json:"is_deleted" db:"is_deleted"`
}

// MealItemUnit 사료/간식 단위 정의 엔티티
type MealItemUnit struct {
	ID           uint         `json:"id" db:"id"`
	MealItemID   uint         `json:"meal_item_id" db:"meal_item_id"`
	MealItemType MealItemType `json:"meal_item_type" db:"meal_item_type"`
	Unit         string       `json:"unit" db:"unit"`
	UnitValue    string       `json:"unit_value" db:"unit_value"`
}

// MealHistory 급여 기록 엔티티
type MealHistory struct {
	ID           uint       `json:"id" db:"id"`
	MealItemID   uint       `json:"meal_item_id" db:"meal_item_id"`
	HistoryDate  time.Time  `json:"history_date" db:"history_date"`
	MealType     MealType   `json:"meal_type" db:"meal_type"`
	Name         string     `json:"name" db:"name"`
	Count        int        `json:"count" db:"count"`
	MealCategory string     `json:"meal_category" db:"meal_category"`
	FeedAt       *time.Time `json:"feed_at" db:"feed_at"`
	ModifiedAt   *time.Time `json:"modified_at" db:"modified_at"`
	DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
	IsDeleted    bool       `json:"is_deleted" db:"is_deleted"`
}

// MealHistoryUnit 급여 기록 단위 엔티티
type MealHistoryUnit struct {
	ID            uint         `json:"id" db:"id"`
	MealHistoryID uint         `json:"meal_histories_id" db:"meal_histories_id"`
	MealItemType  MealItemType `json:"meal_item_type" db:"meal_item_type"`
	Unit          string       `json:"unit" db:"unit"`
	UnitValue     string       `json:"unit_value" db:"unit_value"`
}

// MealItemCreateRequest 사료/간식 아이템 생성 요청
type MealItemCreateRequest struct {
	PetUserRoleID uint     `json:"pet_user_role_id" validate:"required"`
	DataType      DataType `json:"data_type" validate:"required,oneof=FIXED VARIATION"`
	MealType      MealType `json:"meal_type" validate:"required"`
	MealCategory  string   `json:"meal_category" validate:"required"`
	Name          string   `json:"name" validate:"required"`
	UnitType      UnitType `json:"unit_type" validate:"required"`
}

// MealItemUpdateRequest 사료/간식 아이템 수정 요청
type MealItemUpdateRequest struct {
	DataType     *DataType `json:"data_type,omitempty"`
	MealType     *MealType `json:"meal_type,omitempty"`
	MealCategory *string   `json:"meal_category,omitempty"`
	Name         *string   `json:"name,omitempty"`
	UnitType     *UnitType `json:"unit_type,omitempty"`
}

// MealItemUnitCreateRequest 사료/간식 단위 생성 요청
type MealItemUnitCreateRequest struct {
	MealItemID   uint         `json:"meal_item_id" validate:"required"`
	MealItemType MealItemType `json:"meal_item_type" validate:"required"`
	Unit         string       `json:"unit" validate:"required"`
	UnitValue    string       `json:"unit_value" validate:"required"`
}

// MealItemUnitUpdateRequest 사료/간식 단위 수정 요청
type MealItemUnitUpdateRequest struct {
	MealItemType *MealItemType `json:"meal_item_type,omitempty"`
	Unit         *string       `json:"unit,omitempty"`
	UnitValue    *string       `json:"unit_value,omitempty"`
}

// MealHistoryCreateRequest 급여 기록 생성 요청
type MealHistoryCreateRequest struct {
	MealItemID   uint       `json:"meal_item_id" validate:"required"`
	HistoryDate  time.Time  `json:"history_date" validate:"required"`
	MealType     MealType   `json:"meal_type" validate:"required"`
	Name         string     `json:"name" validate:"required"`
	Count        int        `json:"count" validate:"min=1"`
	MealCategory string     `json:"meal_category" validate:"required"`
	FeedAt       *time.Time `json:"feed_at,omitempty"`
}

// MealHistoryUpdateRequest 급여 기록 수정 요청
type MealHistoryUpdateRequest struct {
	HistoryDate  *time.Time `json:"history_date,omitempty"`
	MealType     *MealType  `json:"meal_type,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Count        *int       `json:"count,omitempty"`
	MealCategory *string    `json:"meal_category,omitempty"`
	FeedAt       *time.Time `json:"feed_at,omitempty"`
}

// MealHistoryUnitCreateRequest 급여 기록 단위 생성 요청
type MealHistoryUnitCreateRequest struct {
	MealHistoryID uint         `json:"meal_histories_id" validate:"required"`
	MealItemType  MealItemType `json:"meal_item_type" validate:"required"`
	Unit          string       `json:"unit" validate:"required"`
	UnitValue     string       `json:"unit_value" validate:"required"`
}

// MealHistoryUnitUpdateRequest 급여 기록 단위 수정 요청
type MealHistoryUnitUpdateRequest struct {
	MealItemType *MealItemType `json:"meal_item_type,omitempty"`
	Unit         *string       `json:"unit,omitempty"`
	UnitValue    *string       `json:"unit_value,omitempty"`
}

// MealItemWithUnits 사료/간식 아이템과 단위 정보를 포함한 구조체
type MealItemWithUnits struct {
	MealItem
	Units []MealItemUnit `json:"units"`
}

// MealHistoryWithUnits 급여 기록과 단위 정보를 포함한 구조체
type MealHistoryWithUnits struct {
	MealHistory
	Units []MealHistoryUnit `json:"units"`
}
