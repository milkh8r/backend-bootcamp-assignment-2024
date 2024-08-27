package domain

import (
	"database/sql"
	"time"
)

type Flat struct {
	ID               int64        `json:"id"` // PK
	Number           string       `json:"number"`
	HouseID          int64        `json:"house_id"` // FK
	Price            int64        `json:"price"`
	RoomCount        int64        `json:"room_count"`
	ModerationStatus string       `json:"moderation_status"` // FK
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        sql.NullTime `json:"updated_at"`
}

type FlatRepository interface {
	Create(flat *Flat) error
	GetByID(id int64) (*Flat, error)
	GetByHouseID(houseID int64, userRole string) ([]*Flat, error)
	Update(flat *Flat) error
}
