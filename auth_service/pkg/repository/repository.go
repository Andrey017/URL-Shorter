package repository

import (
	"auth_service"
	"database/sql"
)

type Authorization interface {
	CreateUser(user auth_service.User) (int, error)
	GetUser(username, password string) (auth_service.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQL(db),
	}
}
