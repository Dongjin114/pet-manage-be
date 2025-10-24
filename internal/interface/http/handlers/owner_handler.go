package handlers

import (
	"net/http"
	"strconv"

	"pet-manage-be/internal/domain/entities"
	"pet-manage-be/internal/usecase/owner"

	"github.com/gin-gonic/gin"
)

// OwnerHandler 소유자 핸들러
type OwnerHandler struct {
	ownerUsecase *owner.OwnerUsecase
}

// NewOwnerHandler 새로운 소유자 핸들러 생성
func NewOwnerHandler(ownerUsecase *owner.OwnerUsecase) *OwnerHandler {
	return &OwnerHandler{
		ownerUsecase: ownerUsecase,
	}
}

// GetAllOwners 모든 소유자 조회
func (h *OwnerHandler) GetAllOwners(c *gin.Context) {
	owners, err := h.ownerUsecase.GetAllOwners(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    owners,
		"count":   len(owners),
	})
}

// GetOwnerByID 특정 소유자 조회
func (h *OwnerHandler) GetOwnerByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 소유자 ID입니다",
		})
		return
	}

	owner, err := h.ownerUsecase.GetOwnerByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    owner,
	})
}

// CreateOwner 새로운 소유자 생성
func (h *OwnerHandler) CreateOwner(c *gin.Context) {
	var req entities.OwnerCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	owner, err := h.ownerUsecase.CreateOwner(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    owner,
		"message": "소유자가 성공적으로 생성되었습니다",
	})
}

// UpdateOwner 소유자 정보 수정
func (h *OwnerHandler) UpdateOwner(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 소유자 ID입니다",
		})
		return
	}

	var req entities.OwnerUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := h.ownerUsecase.UpdateOwner(c.Request.Context(), uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "소유자 정보가 성공적으로 수정되었습니다",
		"id":      id,
	})
}

// DeleteOwner 소유자 삭제
func (h *OwnerHandler) DeleteOwner(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 소유자 ID입니다",
		})
		return
	}

	if err := h.ownerUsecase.DeleteOwner(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "소유자가 성공적으로 삭제되었습니다",
		"id":      id,
	})
}
