package repository

import (
	"GinEchoCrud/internal/database"
	"GinEchoCrud/internal/interfaces"
	"GinEchoCrud/internal/models"
	"database/sql"
)

const (
	getUserByEmail = `SELECT id, name, age, email, password FROM users WHERE email = ?`
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository() interfaces.AuthRepositoryInterface {
	return &authRepository{
		db: database.DB,
	}
}

func (r *authRepository) Login(email string) (*models.User, error) {
	row := r.db.QueryRow(getUserByEmail, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
