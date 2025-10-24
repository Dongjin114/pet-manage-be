package repositories

import (
	"context"
	"pet-manage-be/internal/domain/entities"
)

// OwnerRepository 소유자 리포지토리 인터페이스
type OwnerRepository interface {
	// GetAll 모든 소유자 조회
	GetAll(ctx context.Context) ([]entities.Owner, error)

	// GetByID ID로 소유자 조회
	GetByID(ctx context.Context, id uint) (*entities.Owner, error)

	// Create 소유자 생성
	Create(ctx context.Context, owner *entities.Owner) error

	// Update 소유자 수정
	Update(ctx context.Context, id uint, owner *entities.Owner) error

	// Delete 소유자 삭제
	Delete(ctx context.Context, id uint) error

	// GetByEmail 이메일로 소유자 조회
	GetByEmail(ctx context.Context, email string) (*entities.Owner, error)
}
