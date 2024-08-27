package postgres

import (
	"avito-backend-bootcamp/internal/domain"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (email, password, role, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(query, user.Email, user.Password, user.Role, user.CreatedAt).Scan(&user.ID)
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, email, password, role, created_at FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByID(id int64) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, email, password, role, created_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
