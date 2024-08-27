package usecase

import (
	"avito-backend-bootcamp/internal/domain"
	"database/sql"
	"errors"
	"time"
)

type HouseUseCase struct {
	houseRepo domain.HouseRepository
}

func NewHouseUseCase(houseRepo domain.HouseRepository) *HouseUseCase {
	return &HouseUseCase{houseRepo: houseRepo}
}

func (uc *HouseUseCase) CreateHouse(house *domain.House, userRole string) error {
	if userRole != "moderator" {
		return errors.New("only moderators can create houses")
	}

	house.CreatedAt = time.Now()
	return uc.houseRepo.Create(house)
}

func (uc *HouseUseCase) GetHouse(id int64) (*domain.House, error) {
	return uc.houseRepo.GetByID(id)
}

func (uc *HouseUseCase) UpdateHouseLastFlatAdded(id int64) error {
	house, err := uc.houseRepo.GetByID(id)
	if err != nil {
		return err
	}
	house.LastFlatAddedAt = sql.NullTime{Time: time.Now(), Valid: true}
	return uc.houseRepo.Update(house)
}

func (uc *HouseUseCase) ListHouses(page, pageSize int) ([]*domain.House, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	return uc.houseRepo.List(pageSize, offset)
}
