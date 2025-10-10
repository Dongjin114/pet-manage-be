package handlers

import (
	"net/http"
	"strconv"

	"pet-manage-be/models"
	"pet-manage-be/services"

	"github.com/gin-gonic/gin"
)

// CreateMealItem creates a new meal item
func CreateMealItem(c *gin.Context) {
	var req struct {
		DataType string  `json:"data_type" binding:"required"`
		Name     string  `json:"name" binding:"required"`
		Amount   float64 `json:"amount"`
		Unit     string  `json:"unit"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// NewMealType으로 생성 시 자동 검증
	mealType, err := models.NewMealType(req.DataType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// 서비스를 통한 데이터베이스 저장
	mealItem := models.MealItem{
		DataType: mealType,
		Name:     req.Name,
		Amount:   req.Amount,
		Unit:     req.Unit,
	}

	mealService := services.NewMealService()
	if err := mealService.CreateMealItem(&mealItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    mealItem,
		"message": "Meal item created successfully",
	})
}

// GetMealTypes returns all available meal types
func GetMealTypes(c *gin.Context) {
	mealTypes := models.GetAllMealTypes()
	
	// 각 타입의 유효성도 확인
	validTypes := make([]models.MealType, 0)
	for _, mealType := range mealTypes {
		if mealType.IsValid() {
			validTypes = append(validTypes, mealType)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    validTypes,
		"count":   len(validTypes),
	})
}

// GetMealItems gets all meal items
func GetMealItems(c *gin.Context) {
	mealService := services.NewMealService()
	mealItems, err := mealService.GetAllMealItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    mealItems,
		"count":   len(mealItems),
	})
}

// GetMealItem gets a specific meal item
func GetMealItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid meal item ID",
		})
		return
	}

	mealService := services.NewMealService()
	mealItem, err := mealService.GetMealItemByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Meal item not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    mealItem,
	})
}

// UpdateMealItem updates an existing meal item
func UpdateMealItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid meal item ID",
		})
		return
	}

	var req struct {
		DataType *string  `json:"data_type,omitempty"`
		Name     *string  `json:"name,omitempty"`
		Amount   *float64 `json:"amount,omitempty"`
		Unit     *string  `json:"unit,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// DataType이 제공된 경우 유효성 검사
	if req.DataType != nil {
		_, err := models.NewMealType(*req.DataType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
	}

	// 업데이트 데이터 준비
	updates := make(map[string]interface{})
	if req.DataType != nil {
		updates["data_type"] = *req.DataType
	}
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Amount != nil {
		updates["amount"] = *req.Amount
	}
	if req.Unit != nil {
		updates["unit"] = *req.Unit
	}

	mealService := services.NewMealService()
	if err := mealService.UpdateMealItem(uint(id), updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Meal item updated successfully",
		"id":      id,
	})
}

// DeleteMealItem deletes a meal item
func DeleteMealItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid meal item ID",
		})
		return
	}

	mealService := services.NewMealService()
	if err := mealService.DeleteMealItem(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Meal item deleted successfully",
		"id":      id,
	})
}
