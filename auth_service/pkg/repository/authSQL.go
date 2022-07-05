package repository

import (
	"auth_service"
	"database/sql"
	"fmt"
)

type AuthSQL struct {
	db *sql.DB
}

func NewAuthSQL(db *sql.DB) *AuthSQL {
	return &AuthSQL{
		db: db,
	}
}

func (r *AuthSQL) CreateUser(user auth_service.User) (int, error) {
	var idResult int64

	query := fmt.Sprintf("INSERT INTO %s (name, username, password) VALUES ('%s', '%s', '%s')", usersTable, user.Name, user.Username, user.Password)

	result, err := r.db.Exec(query)

	if err != nil {
		return 0, err
	}

	idResult, err = result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(idResult), nil
}

func (r *AuthSQL) GetUser(username, password string) (auth_service.User, error) {
	var user auth_service.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username='%s' AND password='%s'", usersTable, username, password)

	rows, err := r.db.Query(query)

	for rows.Next() {
		rows.Scan(&user.Id)
	}

	return user, err
}
