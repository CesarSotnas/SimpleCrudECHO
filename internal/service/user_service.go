package service

import (
	"GinEchoCrud/internal/interfaces"
	"GinEchoCrud/internal/models"
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
