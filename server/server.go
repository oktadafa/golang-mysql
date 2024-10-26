package server

import "database/sql"

type Server struct {
	sql *sql.DB
}

func NewServer() (*sql.DB, error) {
	var err error
	db, err := sql.Open("mysql", "daffa:okta54321@tcp(localhost:3306)/golang_mysql")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
