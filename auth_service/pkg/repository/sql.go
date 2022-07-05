package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const usersTable = "user"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewSQLDB(cnf Config) (*sql.DB, error) {
	db, errCreateConnDB := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cnf.Username, cnf.Password, cnf.Host, cnf.Port, cnf.DBName))

	if errCreateConnDB != nil {
		return nil, errCreateConnDB
	}

	errPingDB := db.Ping()

	if errPingDB != nil {
		return nil, errPingDB
	}

	return db, nil
}
