package services

import (
	"ColombiaCar/internal/models"
	"ColombiaCar/internal/repository"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, email, password string) (string, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepository: userRepo}
}

func (s *authService) Register(ctx context.Context, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepository.CreateUser(ctx, user)
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
