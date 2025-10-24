package entities

import (
	"time"
)

// Pet 펫 엔티티
type Pet struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Species     string    `json:"species"`
	Breed       string    `json:"breed"`
	Age         int       `json:"age"`
	Gender      string    `json:"gender"`
	Color       string    `json:"color"`
	Weight      float64   `json:"weight"`
	MicrochipID string    `json:"microchip_id"`
	OwnerID     uint      `json:"owner_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PetCreateRequest 펫 생성 요청
type PetCreateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Species     string  `json:"species" validate:"required"`
	Breed       string  `json:"breed"`
	Age         int     `json:"age"`
	Gender      string  `json:"gender"`
	Color       string  `json:"color"`
	Weight      float64 `json:"weight"`
	MicrochipID string  `json:"microchip_id"`
	OwnerID     uint    `json:"owner_id" validate:"required"`
}

// PetUpdateRequest 펫 수정 요청
type PetUpdateRequest struct {
	Name        *string  `json:"name,omitempty"`
	Species     *string  `json:"species,omitempty"`
	Breed       *string  `json:"breed,omitempty"`
	Age         *int     `json:"age,omitempty"`
	Gender      *string  `json:"gender,omitempty"`
	Color       *string  `json:"color,omitempty"`
	Weight      *float64 `json:"weight,omitempty"`
	MicrochipID *string  `json:"microchip_id,omitempty"`
	OwnerID     *uint    `json:"owner_id,omitempty"`
}
