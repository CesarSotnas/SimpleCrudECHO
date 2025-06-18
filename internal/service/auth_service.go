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

func (s *authService) Login(email, password string) (*models.User, error) {
	user, err := s.authRepository.Login(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	return user, nil
}
