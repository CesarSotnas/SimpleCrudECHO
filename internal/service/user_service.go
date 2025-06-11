package service

import (
	"GinEchoCrud/internal/interfaces"
	"GinEchoCrud/internal/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository interfaces.UserRepositoryInterface
}

func NewUserService(userRepository interfaces.UserRepositoryInterface) interfaces.UserServiceInterface {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	return user, nil
}
