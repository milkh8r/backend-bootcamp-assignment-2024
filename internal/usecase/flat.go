package usecase

import (
	"avito-backend-bootcamp/internal/domain"
	"database/sql"
	"errors"
	"time"
)

type FlatUseCase struct {
	flatRepo  domain.FlatRepository
	houseRepo domain.HouseRepository
}

func NewFlatUseCase(flatRepo domain.FlatRepository, houseRepo domain.HouseRepository) *FlatUseCase {
	return &FlatUseCase{
		flatRepo:  flatRepo,
		houseRepo: houseRepo,
	}
}

func (uc *FlatUseCase) CreateFlat(flat *domain.Flat) error {
	flat.ModerationStatus = "created"
	flat.CreatedAt = time.Now()
	flat.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	err := uc.flatRepo.Create(flat)
	if err != nil {
		return err
	}

	house, err := uc.houseRepo.GetByID(flat.HouseID)
	if err != nil {
		return err
	}
	house.LastFlatAddedAt = sql.NullTime{Time: time.Now(), Valid: true}
	return uc.houseRepo.Update(house)
}

func (uc *FlatUseCase) UpdateFlatModerationStatus(id int64, status string, userRole string) error {
	if userRole != "moderator" {
		return errors.New("only moderators can update moderation status")
	}

	flat, err := uc.flatRepo.GetByID(id)
	if err != nil {
		return err
	}

	flat.ModerationStatus = status
	flat.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return uc.flatRepo.Update(flat)
}

func (uc *FlatUseCase) GetFlatsByHouseID(houseID int64, userRole string) ([]*domain.Flat, error) {
	return uc.flatRepo.GetByHouseID(houseID, userRole)
}
