package main

import (
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindById(userId int) (*User, error) {
	query := `
		select id, name, email
		from users
		where id = $1
	`
	var user User
	err := r.DB.QueryRow(query, userId).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(userEmail string) (*User, error) {
	query := `
		select id, name, email, password
		from users
		where email = $1
	`
	var user User
	err := r.DB.QueryRow(query, userEmail).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
