package repository

import (
	"context"
	"errors"
	"pet-manage-be/internal/domain/entities"
	"pet-manage-be/internal/domain/repositories"
)

// ownerRepositoryImpl 소유자 리포지토리 구현체
type ownerRepositoryImpl struct {
	// TODO: 데이터베이스 연결 추가
}

// NewOwnerRepository 새로운 소유자 리포지토리 생성
func NewOwnerRepository() repositories.OwnerRepository {
	return &ownerRepositoryImpl{}
}

// GetAll 모든 소유자 조회
func (r *ownerRepositoryImpl) GetAll(ctx context.Context) ([]entities.Owner, error) {
	// TODO: 데이터베이스에서 조회
	// 현재는 목 데이터 반환
	owners := []entities.Owner{
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
	return owners, nil
}

// GetByID ID로 소유자 조회
func (r *ownerRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.Owner, error) {
	// TODO: 데이터베이스에서 조회
	// 현재는 목 데이터 반환
	owner := &entities.Owner{
		ID:        id,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     "555-0123",
		Address:   "123 Main St",
		City:      "Anytown",
		State:     "CA",
		ZipCode:   "12345",
	}
	return owner, nil
}

// Create 소유자 생성
func (r *ownerRepositoryImpl) Create(ctx context.Context, owner *entities.Owner) error {
	// TODO: 데이터베이스에 저장
	// 현재는 성공으로 처리
	return nil
}

// Update 소유자 수정
func (r *ownerRepositoryImpl) Update(ctx context.Context, id uint, owner *entities.Owner) error {
	// TODO: 데이터베이스에서 수정
	// 현재는 성공으로 처리
	return nil
}

// Delete 소유자 삭제
func (r *ownerRepositoryImpl) Delete(ctx context.Context, id uint) error {
	// TODO: 데이터베이스에서 삭제
	// 현재는 성공으로 처리
	return nil
}

// GetByEmail 이메일로 소유자 조회
func (r *ownerRepositoryImpl) GetByEmail(ctx context.Context, email string) (*entities.Owner, error) {
	// TODO: 데이터베이스에서 조회
	// 현재는 존재하지 않음으로 처리
	return nil, errors.New("소유자를 찾을 수 없습니다")
}
