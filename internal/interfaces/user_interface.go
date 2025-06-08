package interfaces

import "GinEchoCrud/internal/models"

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
}

type UserRepositoryInterface interface {
	GetAllUsers() ([]models.User, error)
}
