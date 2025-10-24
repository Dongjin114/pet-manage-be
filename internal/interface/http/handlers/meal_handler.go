package handlers

import (
	"net/http"
	"strconv"

	"pet-manage-be/internal/domain/entities"
	"pet-manage-be/internal/usecase/meal"

	"github.com/gin-gonic/gin"
)

// MealHandler 급식 핸들러
type MealHandler struct {
	mealUsecase *meal.MealUsecase
}

// NewMealHandler 새로운 급식 핸들러 생성
func NewMealHandler(mealUsecase *meal.MealUsecase) *MealHandler {
	return &MealHandler{
		mealUsecase: mealUsecase,
	}
}

// GetAllMealItems 모든 급식 아이템 조회
func (h *MealHandler) GetAllMealItems(c *gin.Context) {
	mealItems, err := h.mealUsecase.GetAllMealItems(c.Request.Context())
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

// GetMealItemByID 특정 급식 아이템 조회
func (h *MealHandler) GetMealItemByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 급식 아이템 ID입니다",
		})
		return
	}

	mealItem, err := h.mealUsecase.GetMealItemByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    mealItem,
	})
}

// CreateMealItem 새로운 급식 아이템 생성
func (h *MealHandler) CreateMealItem(c *gin.Context) {
	var req entities.MealItemCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	mealItem, err := h.mealUsecase.CreateMealItem(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    mealItem,
		"message": "급식 아이템이 성공적으로 생성되었습니다",
	})
}

// UpdateMealItem 급식 아이템 수정
func (h *MealHandler) UpdateMealItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 급식 아이템 ID입니다",
		})
		return
	}

	var req entities.MealItemUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := h.mealUsecase.UpdateMealItem(c.Request.Context(), uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "급식 아이템이 성공적으로 수정되었습니다",
		"id":      id,
	})
}

// DeleteMealItem 급식 아이템 삭제
func (h *MealHandler) DeleteMealItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 급식 아이템 ID입니다",
		})
		return
	}

	if err := h.mealUsecase.DeleteMealItem(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "급식 아이템이 성공적으로 삭제되었습니다",
		"id":      id,
	})
}

// GetMealTypes 급식 타입 목록 조회
func (h *MealHandler) GetMealTypes(c *gin.Context) {
	mealTypes := h.mealUsecase.GetMealTypes(c.Request.Context())

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    mealTypes,
		"count":   len(mealTypes),
	})
}

// GetMealItemsByType 타입별 급식 아이템 조회
func (h *MealHandler) GetMealItemsByType(c *gin.Context) {
	typeStr := c.Param("type")
	var mealType entities.MealType

	switch typeStr {
	case "FIXED":
		mealType = entities.MealTypeFixed
	case "VARIATION":
		mealType = entities.MealTypeVariation
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 급식 타입입니다",
		})
		return
	}

	mealItems, err := h.mealUsecase.GetMealItemsByType(c.Request.Context(), mealType)
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
