package interfaces

import "GinEchoCrud/internal/models"

type AuthServiceInterface interface {
	Login(email string, password string) (*models.User, error)
}

type AuthRepositoryInterface interface {
	Login(email string) (*models.User, error)
}
