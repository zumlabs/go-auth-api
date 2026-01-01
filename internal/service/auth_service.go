package service

import (
	"errors"

	"github.com/zumlabs/go-auth-api/internal/model"
	"github.com/zumlabs/go-auth-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Register(name, email, password string) error {
	// cek email sudah ada
	_, err := s.UserRepo.FindByEmail(email)
	if err == nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.UserRepo.Create(user)
}
