package repository

import (
	"GinEchoCrud/internal/database"
	"GinEchoCrud/internal/interfaces"
	"GinEchoCrud/internal/models"
	"database/sql"
)

const (
	retrieveAllUsers    = `SELECT id, name, age, email FROM users`
	retrieveUserByEmail = `SELECT id, name, age, email, password FROM users WHERE email = ?`
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() interfaces.UserRepositoryInterface {
	return &userRepository{
		db: database.DB,
	}
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.db.Query(retrieveAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	row := r.db.QueryRow(retrieveUserByEmail, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
