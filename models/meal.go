package models

import (
	"time"
)

// MealItem represents a meal item in the system
type MealItem struct {
	ID        uint      `json:"id"`
	DataType  MealType  `json:"data_type"`
	Name      string    `json:"name" gorm:"not null"`
	Amount    float64   `json:"amount"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MealHistory represents meal history
type MealHistory struct {
	ID        uint      `json:"id" `
	PetID     uint      `json:"pet_id" `
	MealID    uint      `json:"meal_id"`
	Amount    float64   `json:"amount"`
	FedAt     time.Time `json:"fed_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MealUnit represents meal unit
type MealUnit struct {
	ID        uint      `json:"id" `
	Name      string    `json:"name" `
	Symbol    string    `json:"symbol"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
