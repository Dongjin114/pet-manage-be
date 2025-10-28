package handlers

import (
	"net/http"
	"strconv"

	"pet-manage-be/internal/domain/dto"
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

// GetAllMealss 모든 급식 아이템 조회
func (h *MealHandler) GetAllMealss(c *gin.Context) {
	Mealss, err := h.mealUsecase.GetAllMealss(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Mealss,
		"count":   len(Mealss),
	})
}

// GetMealsByID 특정 급식 아이템 조회
func (h *MealHandler) GetMealsByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 급식 아이템 ID입니다",
		})
		return
	}

	Meals, err := h.mealUsecase.GetMealsByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Meals,
	})
}

// CreateMeals 새로운 급식 아이템 생성
func (h *MealHandler) CreateMeals(c *gin.Context) {
	var req dto.MealsCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	Meals, err := h.mealUsecase.CreateMeals(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    Meals,
		"message": "급식 아이템이 성공적으로 생성되었습니다",
	})
}

// UpdateMeals 급식 아이템 수정
func (h *MealHandler) UpdateMeals(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 급식 아이템 ID입니다",
		})
		return
	}

	var req dto.MealsUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := h.mealUsecase.UpdateMeals(c.Request.Context(), uint(id), req); err != nil {
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

// DeleteMeals 급식 아이템 삭제
func (h *MealHandler) DeleteMeals(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 급식 아이템 ID입니다",
		})
		return
	}

	if err := h.mealUsecase.DeleteMeals(c.Request.Context(), uint(id)); err != nil {
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

// GetMealssByType 타입별 급식 아이템 조회
func (h *MealHandler) GetMealssByType(c *gin.Context) {
	typeStr := c.Param("type")
	var mealType entities.MealType

	switch typeStr {
	case "사료":
		mealType = entities.MealTypeFeed
	case "간식":
		mealType = entities.MealTypeSnack
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 급식 타입입니다",
		})
		return
	}

	Mealss, err := h.mealUsecase.GetMealssByType(c.Request.Context(), mealType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Mealss,
		"count":   len(Mealss),
	})
}
