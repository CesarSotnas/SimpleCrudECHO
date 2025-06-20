package interfaces

import (
	"GinEchoCrud/internal/models"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, int, error)
	GetUsersByID(ID int) (models.User, int, error)
	CreateUser(requestUser models.User) (models.User, int, error)
}

type UserRepositoryInterface interface {
	GetAllUsers() ([]models.User, int, error)
	GetUsersByID(ID int) (models.User, int, error)
	CreateUser(requestUser models.User) (models.User, int, error)
}
