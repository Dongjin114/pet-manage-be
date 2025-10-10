package models

import "errors"

// MealType represents the type of meal
type MealType int

// 사료및 간식 목록(FIXED), 급여 정보 타입(VARIATION)
const (
	MealTypeFixed MealType = iota
	MealTypeVariation
	NotValidMealType
)

// String returns the string representation
func (mt MealType) String() string {
	switch mt {
	case MealTypeFixed:
		return "FIXED"
	case MealTypeVariation:
		return "VARIATION"
	default:
		return "UNKNOWN"
	}
}

// IsValid checks if the meal type is valid
func (mt MealType) IsValid() bool {
	return mt == MealTypeFixed || mt == MealTypeVariation
}

// StringToMealType converts string to MealType
func StringToMealType(value string) (MealType, error) {
	switch value {
	case "FIXED":
		return MealTypeFixed, nil
	case "VARIATION":
		return MealTypeVariation, nil
	default:
		return NotValidMealType, errors.New("invalid meal type: must be FIXED or VARIATION")
	}
}
