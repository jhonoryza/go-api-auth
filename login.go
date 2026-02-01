package main

import (
	"database/sql"
	"errors"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=3,max=255"`
}

func Login(db *sql.DB, req LoginRequest) (string, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return "", err
	}

	userRepo := NewUserRepository(db)
	user, err := userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	token, err := GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", errors.New("Generating token failed")
	}

	return token, nil
}
