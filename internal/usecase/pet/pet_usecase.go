package pet

import (
	"context"
	"errors"
	"pet-manage-be/internal/domain/entities"
	"pet-manage-be/internal/domain/repositories"
)

// PetUsecase 펫 유스케이스
type PetUsecase struct {
	petRepo   repositories.PetRepository
	ownerRepo repositories.OwnerRepository
}

// NewPetUsecase 새로운 펫 유스케이스 생성
func NewPetUsecase(petRepo repositories.PetRepository, ownerRepo repositories.OwnerRepository) *PetUsecase {
	return &PetUsecase{
		petRepo:   petRepo,
		ownerRepo: ownerRepo,
	}
}

// GetAllPets 모든 펫 조회
func (u *PetUsecase) GetAllPets(ctx context.Context) ([]entities.Pet, error) {
	return u.petRepo.GetAll(ctx)
}

// GetPetByID ID로 펫 조회
func (u *PetUsecase) GetPetByID(ctx context.Context, id uint) (*entities.Pet, error) {
	if id == 0 {
		return nil, errors.New("펫 ID는 필수입니다")
	}

	return u.petRepo.GetByID(ctx, id)
}

// CreatePet 새로운 펫 생성
func (u *PetUsecase) CreatePet(ctx context.Context, req entities.PetCreateRequest) (*entities.Pet, error) {
	// 소유자 존재 여부 확인
	owner, err := u.ownerRepo.GetByID(ctx, req.OwnerID)
	if err != nil {
		return nil, errors.New("존재하지 않는 소유자입니다")
	}
	if owner == nil {
		return nil, errors.New("존재하지 않는 소유자입니다")
	}

	// 펫 생성
	pet := &entities.Pet{
		Name:        req.Name,
		Species:     req.Species,
		Breed:       req.Breed,
		Age:         req.Age,
		Gender:      req.Gender,
		Color:       req.Color,
		Weight:      req.Weight,
		MicrochipID: req.MicrochipID,
		OwnerID:     req.OwnerID,
		Status:      "ACTIVE",
	}

	if err := u.petRepo.Create(ctx, pet); err != nil {
		return nil, err
	}

	return pet, nil
}

// UpdatePet 펫 정보 수정
func (u *PetUsecase) UpdatePet(ctx context.Context, id uint, req entities.PetUpdateRequest) error {
	if id == 0 {
		return errors.New("펫 ID는 필수입니다")
	}

	// 기존 펫 조회
	pet, err := u.petRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// 수정할 필드만 업데이트
	if req.Name != nil {
		pet.Name = *req.Name
	}
	if req.Species != nil {
		pet.Species = *req.Species
	}
	if req.Breed != nil {
		pet.Breed = *req.Breed
	}
	if req.Age != nil {
		pet.Age = *req.Age
	}
	if req.Gender != nil {
		pet.Gender = *req.Gender
	}
	if req.Color != nil {
		pet.Color = *req.Color
	}
	if req.Weight != nil {
		pet.Weight = *req.Weight
	}
	if req.MicrochipID != nil {
		pet.MicrochipID = *req.MicrochipID
	}
	if req.OwnerID != nil {
		// 소유자 존재 여부 확인
		owner, err := u.ownerRepo.GetByID(ctx, *req.OwnerID)
		if err != nil {
			return errors.New("존재하지 않는 소유자입니다")
		}
		if owner == nil {
			return errors.New("존재하지 않는 소유자입니다")
		}
		pet.OwnerID = *req.OwnerID
	}

	return u.petRepo.Update(ctx, id, pet)
}

// DeletePet 펫 삭제
func (u *PetUsecase) DeletePet(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("펫 ID는 필수입니다")
	}

	return u.petRepo.Delete(ctx, id)
}

// GetPetsByOwnerID 소유자 ID로 펫 목록 조회
func (u *PetUsecase) GetPetsByOwnerID(ctx context.Context, ownerID uint) ([]entities.Pet, error) {
	if ownerID == 0 {
		return nil, errors.New("소유자 ID는 필수입니다")
	}

	return u.petRepo.GetByOwnerID(ctx, ownerID)
}
