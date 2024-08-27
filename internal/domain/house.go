package domain

import (
	"database/sql"
	"time"
)

type House struct {
	ID              int64          `json:"id"`     // PK
	Number          string         `json:"number"` // Unique
	Address         string         `json:"address"`
	BuildYear       sql.NullInt32  `json:"build_year"` // could be unknown?
	Developer       sql.NullString `json:"developer"`  // FK
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       sql.NullTime   `json:"updated_at"`
	LastFlatAddedAt sql.NullTime   `json:"last_flat_added_at"`
}

type HouseRepository interface {
	Create(house *House) error
	GetByID(id int64) (*House, error)
	Update(house *House) error
	List(limit, offset int) ([]*House, error)
}
