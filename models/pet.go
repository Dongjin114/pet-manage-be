package models

import (
	"time"
)

// Pet represents a pet in the system
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

// PetCreateRequest represents the request body for creating a pet
type PetCreateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Species     string  `json:"species" binding:"required"`
	Breed       string  `json:"breed"`
	Age         int     `json:"age"`
	Gender      string  `json:"gender"`
	Color       string  `json:"color"`
	Weight      float64 `json:"weight"`
	MicrochipID string  `json:"microchip_id"`
	OwnerID     uint    `json:"owner_id" binding:"required"`
}

// PetUpdateRequest represents the request body for updating a pet
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
