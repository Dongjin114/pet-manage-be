package handlers

import (
	"net/http"
	"strconv"

	"pet-manage-be/models"

	"github.com/gin-gonic/gin"
)

// GetOwners retrieves all owners
func GetOwners(c *gin.Context) {
	// TODO: Implement database query
	// For now, return mock data
	owners := []models.Owner{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Phone:     "555-0123",
			Address:   "123 Main St",
			City:      "Anytown",
			State:     "CA",
			ZipCode:   "12345",
		},
		{
			ID:        2,
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane.smith@example.com",
			Phone:     "555-0456",
			Address:   "456 Oak Ave",
			City:      "Somewhere",
			State:     "NY",
			ZipCode:   "67890",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    owners,
		"count":   len(owners),
	})
}

// GetOwner retrieves a specific owner by ID
func GetOwner(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid owner ID",
		})
		return
	}

	// TODO: Implement database query
	// For now, return mock data
	owner := models.Owner{
		ID:        uint(id),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     "555-0123",
		Address:   "123 Main St",
		City:      "Anytown",
		State:     "CA",
		ZipCode:   "12345",
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    owner,
	})
}

// CreateOwner creates a new owner
func CreateOwner(c *gin.Context) {
	var req models.OwnerCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// TODO: Implement database insert
	// For now, return success with the created owner
	owner := models.Owner{
		ID:        1,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Address:   req.Address,
		City:      req.City,
		State:     req.State,
		ZipCode:   req.ZipCode,
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    owner,
		"message": "Owner created successfully",
	})
}

// UpdateOwner updates an existing owner
func UpdateOwner(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid owner ID",
		})
		return
	}

	var req models.OwnerUpdateRequest
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
		"message": "Owner updated successfully",
		"id":      id,
	})
}

// DeleteOwner deletes an owner
func DeleteOwner(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid owner ID",
		})
		return
	}

	// TODO: Implement database delete
	// For now, return success
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Owner deleted successfully",
		"id":      id,
	})
}
