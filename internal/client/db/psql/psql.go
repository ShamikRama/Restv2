package psql

import "database/sql"

type DB struct {
	db *sql.DB
}

func New(db *sql.DB) (*DB, error) {

}
