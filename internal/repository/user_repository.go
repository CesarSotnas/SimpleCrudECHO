package repository

import (
	"GinEchoCrud/internal/database"
	"GinEchoCrud/internal/helpers"
	"GinEchoCrud/internal/interfaces"
	"GinEchoCrud/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

const (
	retrieveAllUsers  = `SELECT id, name, age, email FROM users`
	retrieveUsersByID = `SELECT id, name, age, email FROM users WHERE id = ?`
	createUser        = `INSERT INTO users (name, age, email) VALUES (?, ?, ?);`
	deleteUser        = `DELETE FROM users WHERE id = ?`
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

func (r *userRepository) CreateUser(requestUser models.User) (models.User, int, error) {
	result, err := r.db.Exec(createUser, requestUser.Name, requestUser.Age, requestUser.Email)
	if err != nil {
		return models.User{}, helpers.StatusBadRequest, err
	}

	userID, errUser := result.LastInsertId()
	if errUser != nil {
		return models.User{}, helpers.StatusBadRequest, errUser
	}

	var user models.User

	errQuery := r.db.QueryRow(retrieveUsersByID, userID).Scan(&user.ID, &user.Name, &user.Age, &user.Email)
	if errQuery != nil {
		return models.User{}, helpers.StatusBadRequest, errQuery
	}

	return user, helpers.StatusOk, nil
}

func (r *userRepository) UpdateUser(userID int, user models.UserRequests) (int, error) {
	var (
		fields []string
		args   []interface{}
	)

	if user.Name != nil {
		fields = append(fields, "name = ?")
		args = append(args, *user.Name)
	}

	if user.Age != nil {
		fields = append(fields, "age = ?")
		args = append(args, *user.Age)
	}

	if user.Email != nil {
		fields = append(fields, "email = ?")
		args = append(args, *user.Email)
	}

	if len(fields) == 0 {
		return helpers.StatusBadRequest, helpers.ErrMsgNoFieldsToUpdate
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = ?", strings.Join(fields, ", "))
	args = append(args, userID)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return helpers.StatusBadRequest, err
	}
	return helpers.StatusUpdated, nil
}

func (r *userRepository) DeleteUser(userID int) (int, error) {
	_, err := r.db.Exec(deleteUser, userID)
	if err != nil {
		return helpers.StatusBadRequest, err
	}
	return helpers.StatusOk, nil
}
