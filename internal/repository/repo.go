package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db  *sql.DB
	dsn string
}

func InitRepo(dsn string) (*Repository, error) {
	return &Repository{dsn: dsn}, nil
}

func (r *Repository) openDb() error {
	db, err := sql.Open("sqlite3", r.dsn)
	if err != nil {
		return err
	}
	r.db = db
	return nil
}

func (r *Repository) closeDb() error {
	return r.db.Close()
}
