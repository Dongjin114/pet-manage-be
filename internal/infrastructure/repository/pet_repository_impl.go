package repository

import (
	"context"
	"pet-manage-be/internal/domain/entities"
	"pet-manage-be/internal/domain/repositories"
)

// petRepositoryImpl 펫 리포지토리 구현체
type petRepositoryImpl struct {
	// TODO: 데이터베이스 연결 추가
}

// NewPetRepository 새로운 펫 리포지토리 생성
func NewPetRepository() repositories.PetRepository {
	return &petRepositoryImpl{}
}

// GetAll 모든 펫 조회
func (r *petRepositoryImpl) GetAll(ctx context.Context) ([]entities.Pet, error) {
	// TODO: 데이터베이스에서 조회
	// 현재는 목 데이터 반환
	pets := []entities.Pet{
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
			Status:  "ACTIVE",
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
			Status:  "ACTIVE",
		},
	}
	return pets, nil
}

// GetByID ID로 펫 조회
func (r *petRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.Pet, error) {
	// TODO: 데이터베이스에서 조회
	// 현재는 목 데이터 반환
	pet := &entities.Pet{
		ID:      id,
		Name:    "Buddy",
		Species: "Dog",
		Breed:   "Golden Retriever",
		Age:     3,
		Gender:  "Male",
		Color:   "Golden",
		Weight:  25.5,
		OwnerID: 1,
		Status:  "ACTIVE",
	}
	return pet, nil
}

// Create 펫 생성
func (r *petRepositoryImpl) Create(ctx context.Context, pet *entities.Pet) error {
	// TODO: 데이터베이스에 저장
	// 현재는 성공으로 처리
	return nil
}

// Update 펫 수정
func (r *petRepositoryImpl) Update(ctx context.Context, id uint, pet *entities.Pet) error {
	// TODO: 데이터베이스에서 수정
	// 현재는 성공으로 처리
	return nil
}

// Delete 펫 삭제
func (r *petRepositoryImpl) Delete(ctx context.Context, id uint) error {
	// TODO: 데이터베이스에서 삭제
	// 현재는 성공으로 처리
	return nil
}

// GetByOwnerID 소유자 ID로 펫 목록 조회
func (r *petRepositoryImpl) GetByOwnerID(ctx context.Context, ownerID uint) ([]entities.Pet, error) {
	// TODO: 데이터베이스에서 조회
	// 현재는 목 데이터 반환
	pets := []entities.Pet{
		{
			ID:      1,
			Name:    "Buddy",
			Species: "Dog",
			Breed:   "Golden Retriever",
			Age:     3,
			Gender:  "Male",
			Color:   "Golden",
			Weight:  25.5,
			OwnerID: ownerID,
			Status:  "ACTIVE",
		},
	}
	return pets, nil
}
