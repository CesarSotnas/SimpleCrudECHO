package interfaces

import "GinEchoCrud/internal/models"

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	Login(email string, password string) (*models.User, error)
}

type UserRepositoryInterface interface {
	GetAllUsers() ([]models.User, error)
	GetByEmail(email string) (*models.User, error)
}
