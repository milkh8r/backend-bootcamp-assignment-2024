package postgres

import (
	"avito-backend-bootcamp/internal/domain"
	"database/sql"
)

type flatRepository struct {
	db *sql.DB
}

func NewFlatRepository(db *sql.DB) domain.FlatRepository {
	return &flatRepository{db: db}
}

func (r *flatRepository) Create(flat *domain.Flat) error {
	query := `INSERT INTO flats (house_id, number, price, room_count, moderation_status, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	return r.db.QueryRow(query, flat.HouseID, flat.Number, flat.Price, flat.RoomCount, flat.ModerationStatus, flat.CreatedAt, flat.UpdatedAt).Scan(&flat.ID)
}

func (r *flatRepository) GetByID(id int64) (*domain.Flat, error) {
	flat := &domain.Flat{}
	query := `SELECT id, house_id, number, price, room_count, moderation_status, created_at, updated_at FROM flats WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&flat.ID, &flat.HouseID, &flat.Number, &flat.Price, &flat.RoomCount, &flat.ModerationStatus, &flat.CreatedAt, &flat.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return flat, nil
}

func (r *flatRepository) GetByHouseID(houseID int64, userRole string) ([]*domain.Flat, error) {
	var query string
	if userRole == "moderator" {
		query = `SELECT id, house_id, number, price, room_count, moderation_status, created_at, updated_at FROM flats WHERE house_id = $1`
	} else {
		query = `SELECT id, house_id, number, price, room_count, moderation_status, created_at, updated_at FROM flats WHERE house_id = $1 AND moderation_status = 'approved'`
	}

	rows, err := r.db.Query(query, houseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flats []*domain.Flat
	for rows.Next() {
		flat := &domain.Flat{}
		err := rows.Scan(&flat.ID, &flat.HouseID, &flat.Number, &flat.Price, &flat.RoomCount, &flat.ModerationStatus, &flat.CreatedAt, &flat.UpdatedAt)
		if err != nil {
			return nil, err
		}
		flats = append(flats, flat)
	}

	return flats, nil
}

func (r *flatRepository) Update(flat *domain.Flat) error {
	query := `UPDATE flats SET house_id = $1, number = $2, price = $3, room_count = $4, moderation_status = $5, updated_at = $6 WHERE id = $7`
	_, err := r.db.Exec(query, flat.HouseID, flat.Number, flat.Price, flat.RoomCount, flat.ModerationStatus, flat.UpdatedAt, flat.ID)
	return err
}
