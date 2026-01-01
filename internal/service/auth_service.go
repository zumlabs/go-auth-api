package service

import (
	"errors"

	"github.com/zumlabs/go-auth-api/internal/model"
	"github.com/zumlabs/go-auth-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo  *repository.UserRepository
	JWTSecret string
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		UserRepo:  userRepo,
		JWTSecret: jwtSecret,
	}
}

func (s *AuthService) Register(name, email, password string) error {
	_, err := s.UserRepo.FindByEmail(email)
	if err == nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	return GenerateToken(user.ID, s.JWTSecret)
}
