package domain

import "time"

type User struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	LastLoginAt time.Time `json:"last_login_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
	GetByID(id int64) (*User, error)
}
