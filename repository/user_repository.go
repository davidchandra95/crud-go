package repository

import (
	"crud-go/model"
	"database/sql"
)

type UserRepository interface {
	Create(user *model.User) error
	GetAll() ([]*model.User, error)
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	query := "INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id"
	return r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
}

func (r *userRepository) GetAll() ([]*model.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *userRepository) GetByID(id int) (*model.User, error) {
	query := "SELECT id, name, email FROM users WHERE id=$1"
	row := r.db.QueryRow(query, id)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email=$1"
	row := r.db.QueryRow(query, email)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(user *model.User) error {
	query := "UPDATE users SET name=$1, email=$2 WHERE id=$3"
	_, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

func (r *userRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.db.Exec(query, id)
	return err
}
