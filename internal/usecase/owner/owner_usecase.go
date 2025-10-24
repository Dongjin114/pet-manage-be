package owner

import (
	"context"
	"errors"
	"pet-manage-be/internal/domain/entities"
	"pet-manage-be/internal/domain/repositories"
)

// OwnerUsecase 소유자 유스케이스
type OwnerUsecase struct {
	ownerRepo repositories.OwnerRepository
}

// NewOwnerUsecase 새로운 소유자 유스케이스 생성
func NewOwnerUsecase(ownerRepo repositories.OwnerRepository) *OwnerUsecase {
	return &OwnerUsecase{
		ownerRepo: ownerRepo,
	}
}

// GetAllOwners 모든 소유자 조회
func (u *OwnerUsecase) GetAllOwners(ctx context.Context) ([]entities.Owner, error) {
	return u.ownerRepo.GetAll(ctx)
}

// GetOwnerByID ID로 소유자 조회
func (u *OwnerUsecase) GetOwnerByID(ctx context.Context, id uint) (*entities.Owner, error) {
	if id == 0 {
		return nil, errors.New("소유자 ID는 필수입니다")
	}

	return u.ownerRepo.GetByID(ctx, id)
}

// CreateOwner 새로운 소유자 생성
func (u *OwnerUsecase) CreateOwner(ctx context.Context, req entities.OwnerCreateRequest) (*entities.Owner, error) {
	// 이메일 중복 검사
	existingOwner, err := u.ownerRepo.GetByEmail(ctx, req.Email)
	if err == nil && existingOwner != nil {
		return nil, errors.New("이미 존재하는 이메일입니다")
	}

	// 소유자 생성
	owner := &entities.Owner{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Address:   req.Address,
		City:      req.City,
		State:     req.State,
		ZipCode:   req.ZipCode,
	}

	if err := u.ownerRepo.Create(ctx, owner); err != nil {
		return nil, err
	}

	return owner, nil
}

// UpdateOwner 소유자 정보 수정
func (u *OwnerUsecase) UpdateOwner(ctx context.Context, id uint, req entities.OwnerUpdateRequest) error {
	if id == 0 {
		return errors.New("소유자 ID는 필수입니다")
	}

	// 기존 소유자 조회
	owner, err := u.ownerRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// 수정할 필드만 업데이트
	if req.FirstName != nil {
		owner.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		owner.LastName = *req.LastName
	}
	if req.Email != nil {
		owner.Email = *req.Email
	}
	if req.Phone != nil {
		owner.Phone = *req.Phone
	}
	if req.Address != nil {
		owner.Address = *req.Address
	}
	if req.City != nil {
		owner.City = *req.City
	}
	if req.State != nil {
		owner.State = *req.State
	}
	if req.ZipCode != nil {
		owner.ZipCode = *req.ZipCode
	}

	return u.ownerRepo.Update(ctx, id, owner)
}

// DeleteOwner 소유자 삭제
func (u *OwnerUsecase) DeleteOwner(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("소유자 ID는 필수입니다")
	}

	return u.ownerRepo.Delete(ctx, id)
}
