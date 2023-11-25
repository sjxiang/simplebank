package db

import (
	"database/sql"
)


type Data struct {
	db *sql.DB
}

// NewData .
func NewData(db *sql.DB) *Data {
	return &Data{db: db}
}
