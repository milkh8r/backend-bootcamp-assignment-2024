package usecase

import (
	"avito-backend-bootcamp/internal/domain"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) Register(email, password, role string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		Email:     email,
		Password:  string(hashedPassword),
		Role:      role,
		CreatedAt: time.Now(),
	}

	return uc.userRepo.Create(user)
}

func (uc *UserUseCase) Login(email, password string) (*domain.User, error) {
	user, err := uc.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
