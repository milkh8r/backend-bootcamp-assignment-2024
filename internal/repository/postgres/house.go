package postgres

import (
	"avito-backend-bootcamp/internal/domain"
	"database/sql"
)

type houseRepository struct {
	db *sql.DB
}

func NewHouseRepository(db *sql.DB) domain.HouseRepository {
	return &houseRepository{db: db}
}

func (r *houseRepository) Create(house *domain.House) error {
	query := `INSERT INTO houses (number, address, build_year, developer, created_at, last_flat_added_at) 
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.db.QueryRow(query, house.Number, house.Address, house.BuildYear, house.Developer, house.CreatedAt, house.LastFlatAddedAt).Scan(&house.ID)
}

func (r *houseRepository) GetByID(id int64) (*domain.House, error) {
	house := &domain.House{}
	query := `SELECT id, number, address, build_year, developer, created_at, last_flat_added_at FROM houses WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&house.ID, &house.Number, &house.Address, &house.BuildYear, &house.Developer, &house.CreatedAt, &house.LastFlatAddedAt)
	if err != nil {
		return nil, err
	}
	return house, nil
}

func (r *houseRepository) Update(house *domain.House) error {
	query := `UPDATE houses SET number = $1, address = $2, build_year = $3, developer = $4, last_flat_added_at = $5 WHERE id = $6`
	_, err := r.db.Exec(query, house.Number, house.Address, house.BuildYear, house.Developer, house.LastFlatAddedAt, house.ID)
	return err
}

func (r *houseRepository) List(limit, offset int) ([]*domain.House, error) {
	query := `SELECT id, number, address, build_year, developer, created_at, last_flat_added_at FROM houses ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var houses []*domain.House
	for rows.Next() {
		house := &domain.House{}
		err := rows.Scan(&house.ID, &house.Number, &house.Address, &house.BuildYear, &house.Developer, &house.CreatedAt, &house.LastFlatAddedAt)
		if err != nil {
			return nil, err
		}
		houses = append(houses, house)
	}

	return houses, nil
}
