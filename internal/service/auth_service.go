package service

import (
	"GinEchoCrud/internal/interfaces"
	"GinEchoCrud/internal/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authRepository interfaces.AuthRepositoryInterface
}

func NewAuthService(authRepository interfaces.AuthRepositoryInterface) interfaces.AuthServiceInterface {
	return &authService{
		authRepository: authRepository,
	}
}

func (s *authService) Login(email string, password string) (*models.Admin, error) {
	admin, err := s.authRepository.Login(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	return admin, nil
}
