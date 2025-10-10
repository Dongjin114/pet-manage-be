package handlers

import (
	"net/http"
	"strconv"

	"pet-manage-be/models"

	"github.com/gin-gonic/gin"
)

// GetPets retrieves all pets
func GetPets(c *gin.Context) {
	// TODO: Implement database query
	// For now, return mock data
	pets := []models.Pet{
		{
			ID:      1,
			Name:    "Buddy",
			Species: "Dog",
			Breed:   "Golden Retriever",
			Age:     3,
			Gender:  "Male",
			Color:   "Golden",
			Weight:  25.5,
			OwnerID: 1,
		},
		{
			ID:      2,
			Name:    "Whiskers",
			Species: "Cat",
			Breed:   "Persian",
			Age:     2,
			Gender:  "Female",
			Color:   "White",
			Weight:  4.2,
			OwnerID: 2,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pets,
		"count":   len(pets),
	})
}

// GetPet retrieves a specific pet by ID
func GetPet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid pet ID",
		})
		return
	}

	// TODO: Implement database query
	// For now, return mock data
	pet := models.Pet{
		ID:       uint(id),
		Name:     "Buddy",
		Species:  "Dog",
		Breed:    "Golden Retriever",
		Age:      3,
		Gender:   "Male",
		Color:    "Golden",
		Weight:   25.5,
		OwnerID:  1,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pet,
	})
}

// CreatePet creates a new pet
func CreatePet(c *gin.Context) {
	var req models.PetCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// TODO: Implement database insert
	// For now, return success with the created pet
	pet := models.Pet{
		ID:        1,
		Name:      req.Name,
		Species:   req.Species,
		Breed:     req.Breed,
		Age:       req.Age,
		Gender:    req.Gender,
		Color:     req.Color,
		Weight:    req.Weight,
		MicrochipID: req.MicrochipID,
		OwnerID:   req.OwnerID,
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    pet,
		"message": "Pet created successfully",
	})
}

// UpdatePet updates an existing pet
func UpdatePet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid pet ID",
		})
		return
	}

	var req models.PetUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// TODO: Implement database update
	// For now, return success
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Pet updated successfully",
		"id":      id,
	})
}

// DeletePet deletes a pet
func DeletePet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid pet ID",
		})
		return
	}

	// TODO: Implement database delete
	// For now, return success
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Pet deleted successfully",
		"id":      id,
	})
}
