package repository

import (
	"GinEchoCrud/internal/interfaces"
	"GinEchoCrud/internal/models"
)


func (r *userRepository) GetAllUsers() []models.User {
	users := []models.User{
		{ID: 15, Name: "Carlos", Age: 230, Email: "emailTest1@gmail.com"},
		{ID: 25, Name: "Cesar", Age: 120, Email: "emailTest2@gmail.com"},
		{ID: 5, Name: "Alberto", Age: 40, Email: "emailTest3@gmail.com"},
		{ID: 4, Name: "Nobrega", Age: 82, Email: "emailTest7@gmail.com"},
		{ID: 3, Name: "Joao", Age: 1000, Email: "emailTest10@gmail.com"},
		{ID: 2, Name: "Manoel", Age: 30, Email: "emailTest100@gmail.com"},
		{ID: 1, Name: "Paulo", Age: 18, Email: "emailTest131@gmail.com"},
	}
	return users
}
