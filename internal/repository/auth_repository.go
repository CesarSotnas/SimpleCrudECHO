package repository

import (
	"GinEchoCrud/internal/database"
	"GinEchoCrud/internal/interfaces"
	"GinEchoCrud/internal/models"
	"database/sql"
)

const (
	getUserByEmail = `SELECT id, email, password FROM admin WHERE email = ?`
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository() interfaces.AuthRepositoryInterface {
	return &authRepository{
		db: database.DB,
	}
}

func (r *authRepository) Login(email string) (*models.Admin, error) {
	row := r.db.QueryRow(getUserByEmail, email)

	var admin models.Admin
	err := row.Scan(&admin.ID, &admin.Email, &admin.Password)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
