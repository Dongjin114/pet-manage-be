package handlers

import (
	"net/http"
	"strconv"

	"pet-manage-be/internal/domain/entities"
	"pet-manage-be/internal/usecase/pet"

	"github.com/gin-gonic/gin"
)

// PetHandler 펫 핸들러
type PetHandler struct {
	petUsecase *pet.PetUsecase
}

// NewPetHandler 새로운 펫 핸들러 생성
func NewPetHandler(petUsecase *pet.PetUsecase) *PetHandler {
	return &PetHandler{
		petUsecase: petUsecase,
	}
}

// GetAllPets 모든 펫 조회
func (h *PetHandler) GetAllPets(c *gin.Context) {
	pets, err := h.petUsecase.GetAllPets(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pets,
		"count":   len(pets),
	})
}

// GetPetByID 특정 펫 조회
func (h *PetHandler) GetPetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 펫 ID입니다",
		})
		return
	}

	pet, err := h.petUsecase.GetPetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pet,
	})
}

// CreatePet 새로운 펫 생성
func (h *PetHandler) CreatePet(c *gin.Context) {
	var req entities.PetCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	pet, err := h.petUsecase.CreatePet(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    pet,
		"message": "펫이 성공적으로 생성되었습니다",
	})
}

// UpdatePet 펫 정보 수정
func (h *PetHandler) UpdatePet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 펫 ID입니다",
		})
		return
	}

	var req entities.PetUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := h.petUsecase.UpdatePet(c.Request.Context(), uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "펫 정보가 성공적으로 수정되었습니다",
		"id":      id,
	})
}

// DeletePet 펫 삭제
func (h *PetHandler) DeletePet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 펫 ID입니다",
		})
		return
	}

	if err := h.petUsecase.DeletePet(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "펫이 성공적으로 삭제되었습니다",
		"id":      id,
	})
}

// GetPetsByOwnerID 소유자 ID로 펫 목록 조회
func (h *PetHandler) GetPetsByOwnerID(c *gin.Context) {
	ownerIDStr := c.Param("owner_id")
	ownerID, err := strconv.ParseUint(ownerIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "잘못된 소유자 ID입니다",
		})
		return
	}

	pets, err := h.petUsecase.GetPetsByOwnerID(c.Request.Context(), uint(ownerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pets,
		"count":   len(pets),
	})
}
