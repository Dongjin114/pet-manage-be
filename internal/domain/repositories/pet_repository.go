package repositories

import (
	"context"
	"pet-manage-be/internal/domain/entities"
)

// PetRepository 펫 리포지토리 인터페이스
type PetRepository interface {
	// GetAll 모든 펫 조회
	GetAll(ctx context.Context) ([]entities.Pet, error)

	// GetByID ID로 펫 조회
	GetByID(ctx context.Context, id uint) (*entities.Pet, error)

	// Create 펫 생성
	Create(ctx context.Context, pet *entities.Pet) error

	// Update 펫 수정
	Update(ctx context.Context, id uint, pet *entities.Pet) error

	// Delete 펫 삭제
	Delete(ctx context.Context, id uint) error

	// GetByOwnerID 소유자 ID로 펫 목록 조회
	GetByOwnerID(ctx context.Context, ownerID uint) ([]entities.Pet, error)
}
