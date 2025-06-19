package repository

import (
	"GinEchoCrud/internal/database"
	"GinEchoCrud/internal/helpers"
	"GinEchoCrud/internal/interfaces"
	"GinEchoCrud/internal/models"
	"database/sql"
)

const (
	retrieveAllUsers  = `SELECT id, name, age, email FROM users`
	retrieveUsersByID = `SELECT id, name, age, email FROM users WHERE id = ?`
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() interfaces.UserRepositoryInterface {
	return &userRepository{
		db: database.DB,
	}
}

func (r *userRepository) GetAllUsers() ([]models.User, int, error) {
	rows, err := r.db.Query(retrieveAllUsers)
	if err != nil {
		return nil, helpers.StatusBadRequest, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Email); err != nil {
			return nil, helpers.StatusBadRequest, err
		}
		users = append(users, u)
	}
	return users, helpers.StatusOk, nil
}

func (r *userRepository) GetUsersByID(ID int) (models.User, int, error) {
	var users models.User

	row := r.db.QueryRow(retrieveUsersByID, ID)
	err := row.Scan(&users.ID, &users.Name, &users.Age, &users.Email)

	if err == sql.ErrNoRows {
		return users, helpers.StatusNotFound, helpers.ErrMsgNotFound
	}

	if err != nil {
		return users, helpers.StatusBadRequest, err
	}

	return users, helpers.StatusOk, nil
}
